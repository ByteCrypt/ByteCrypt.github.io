package services

import (
	"bytecrypt_api/database"
	"bytecrypt_api/utils"
	"bytecrypt_api/v1/models"
	"context"
	"errors"
	"fmt"
)

func NewAddAdminParams(id, email, name, username, password string, role utils.Role) (database.AddAdminParams, error) {
	admin := database.AddAdminParams{
		ID:       id,
		Email:    email,
		Name:     name,
		Username: username,
		Password: password,
		Role:     int32(role),
	}

	if err := ValidateAdminParams(&admin); err != nil {
		return database.AddAdminParams{}, err
	}

	return admin, nil
}

func ValidateAdminParams(admin *database.AddAdminParams) error {
	if admin.ID == "" {
		return errors.New("cannot provide an empty id")
	}
	if admin.Email == "" {
		return errors.New("cannot provide an empty email")
	}
	if admin.Name == "" {
		return errors.New("cannot provide an empty name")
	}
	if admin.Username == "" {
		return errors.New("cannot provide an empty username")
	}
	if admin.Password == "" {
		return errors.New("cannot provide an empty password")
	}
	if utils.Role(admin.Role) == utils.InvalidRole {
		return errors.New("cannot provide an erroneous role")
	}

	return nil
}

func ConvertAdminToDb(admin models.Administrator) (database.AddAdminParams, error) {
	dbadmin := database.AddAdminParams{
		ID:       admin.Id,
		Email:    admin.Email,
		Name:     admin.Name,
		Username: admin.Username,
		Password: admin.Password,
		Role:     int32(admin.Role),
	}

	if err := ValidateAdminParams(&dbadmin); err != nil {
		return database.AddAdminParams{}, err
	}

	return dbadmin, nil
}

func ConvertDbToAdmin(dbadmin database.Administrator) (models.Administrator, error) {
	admin := models.Administrator{
		Id:       dbadmin.ID,
		Email:    dbadmin.Email,
		Name:     dbadmin.Name,
		Username: dbadmin.Username,
		Password: dbadmin.Password,
		Role:     utils.Role(dbadmin.Role),
	}

	if err := models.ValidateAdmin(&admin); err != nil {
		return models.BlankAdmin(), err
	}

	return admin, nil
}

func (provider *Provider) AddAdmin(backend utils.Backend, admin *models.Administrator) error {
	// Validate passed email
	if err := provider.ValidateEmail(admin.Email); err != nil {
		return err
	}

	// Validate passed admin struct
	if err := models.ValidateAdmin(admin); err != nil {
		return err
	}

	// Verify uniqueness
	if dbadmin, err := provider.GetAdminByEmail(admin.Email); err == nil {
		if dbadmin.Id == admin.Id {
			err = fmt.Errorf("%w; %s", err, fmt.Sprintf("an administrator already exists with the id: %s", admin.Id))
		}
		err = fmt.Errorf("%w; %s", err, fmt.Sprintf("an administrator already exists with the email: %s", admin.Email))
		if dbadmin.Username == admin.Username {
			err = fmt.Errorf("%w; %s", err, fmt.Sprintf("an administrator already exists with the username: %s", admin.Username))
		}

		return fmt.Errorf("%s", err.Error()[2:])
	}

	// Convert to proper struct
	add, err := ConvertAdminToDb(*admin)
	if err != nil {
		return err
	}

	// Add to db
	dbadmin, err := provider.Queries.AddAdmin(context.Background(), add)
	if err != nil {
		backend.Log <- utils.NewLog(utils.Warn, fmt.Sprintf("Could not add new admin: %s", err.Error()))
		return err
	}
	*admin, err = ConvertDbToAdmin(dbadmin)
	if err != nil {
		return err
	}
	backend.Log <- utils.NewLog(utils.Info, fmt.Sprintf("New administrator added: %v", admin.ToString()))

	return nil
}

func (provider *Provider) RemoveAdmin(backend utils.Backend, admin *models.Administrator) error {
	// Verify admin exists
	if _, err := provider.GetAdminById(admin.Id); err == nil {
		return fmt.Errorf("an administrator with that id does not exist")
	}

	// Delet
	if err := provider.Queries.DeleteAdmin(context.Background(), admin.Id); err != nil {
		return err
	}
	backend.Log <- utils.NewLog(utils.Info, fmt.Sprintf("Admin removed: %v", admin))

	return nil
}

func (provider *Provider) GetAllAdmins() ([]models.Administrator, error) {
	// Get all the admins
	dbadmins, err := provider.Queries.GetAdmins(context.Background())
	if err != nil {
		return nil, err
	}

	// Convert and return
	var admins []models.Administrator
	for _, dba := range dbadmins {
		a, err := ConvertDbToAdmin(dba)
		if err != nil {
			return nil, err
		}
		admins = append(admins, a)
	}

	return admins, nil
}

func (provider *Provider) GetAdminById(id string) (models.Administrator, error) {
	// Validate passed id
	if id == "" {
		return models.BlankAdmin(), fmt.Errorf("id cannot be blank")
	}

	// Get dbadmin
	dbadmin, err := provider.Queries.GetAdminById(context.Background(), id)
	if err != nil {
		return models.BlankAdmin(), err
	}

	return ConvertDbToAdmin(dbadmin)
}

func (provider *Provider) GetAdminByEmail(email string) (models.Administrator, error) {
	//Validate email
	if err := provider.ValidateEmail(email); err != nil {
		return models.BlankAdmin(), err
	}

	// Get dbadmin
	dbadmin, err := provider.Queries.GetAdminByEmail(context.Background(), email)
	if err != nil {
		return models.BlankAdmin(), err
	}

	return ConvertDbToAdmin(dbadmin)
}

func (provider *Provider) GetAdminByUsername(username string) (models.Administrator, error) {
	// Validate username
	if username == "" {
		return models.BlankAdmin(), fmt.Errorf("username cannot be blank")
	}

	// Get dbadmin
	dbadmin, err := provider.Queries.GetAdminByUsername(context.Background(), username)
	if err != nil {
		return models.BlankAdmin(), err
	}

	return ConvertDbToAdmin(dbadmin)
}

func (provider *Provider) UpdateAdminEmail(backend *utils.Backend, admin *models.Administrator, email string) error {
	// Validate email
	if err := provider.ValidateEmail(email); err != nil {
		return err
	}

	// Verify the administrator exists
	if _, err := provider.GetAdminById(admin.Id); err != nil {
		return err
	}

	// Update the admin
	update := database.UpdateAdminEmailParams{
		ID:    admin.Id,
		Email: email,
	}
	dbadmin, err := provider.Queries.UpdateAdminEmail(context.Background(), update)
	if err != nil {
		backend.Log <- utils.NewLog(utils.Warn, fmt.Sprintf("Unable to update administrator <%s> email: %s", admin.Id, err.Error()))
		return err
	}

	// Convert
	a, err := ConvertDbToAdmin(dbadmin)
	if err != nil {
		return err
	}
	*admin = a
	backend.Log <- utils.NewLog(utils.Info, fmt.Sprintf("Updated administrator email: %s", admin.ToString()))

	return nil
}

func (provider *Provider) UpdateAdminName(backend *utils.Backend, admin *models.Administrator, name string) error {
	// Validate the name
	if name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	// Verify the admin exists
	if _, err := provider.GetAdminById(admin.Id); err != nil {
		return err
	}

	// Updat ethe admin
	update := database.UpdateAdminNameParams{
		ID:   admin.Id,
		Name: name,
	}
	dbadmin, err := provider.Queries.UpdateAdminName(context.Background(), update)
	if err != nil {
		backend.Log <- utils.NewLog(utils.Warn, fmt.Sprintf("Unable to update administrator <%s> name: %s", admin.Id, err.Error()))
		return err
	}

	// Convert and return
	a, err := ConvertDbToAdmin(dbadmin)
	if err != nil {
		return err
	}
	*admin = a
	backend.Log <- utils.NewLog(utils.Info, fmt.Sprintf("Updated administrator name: %v", admin.ToString()))

	return nil
}

func (provider *Provider) UpdateAdminUsername(backend *utils.Backend, admin *models.Administrator, username string) error {
	// Validate username
	if len(username) == 0 {
		return fmt.Errorf("username cannot be empty")
	}

	// Verify the admin exists
	if _, err := provider.GetAdminById(admin.Id); err != nil {
		return err
	}

	// Update the username
	update := database.UpdateAdminUsernameParams{
		ID:       admin.Id,
		Username: username,
	}
	dbadmin, err := provider.Queries.UpdateAdminUsername(context.Background(), update)
	if err != nil {
		backend.Log <- utils.NewLog(utils.Warn, fmt.Sprintf("Unable to update administrator <%s> username: %s", admin.Id, err.Error()))
		return err
	}
	a, err := ConvertDbToAdmin(dbadmin)
	if err != nil {
		return err
	}
	*admin = a
	backend.Log <- utils.NewLog(utils.Info, fmt.Sprintf("Updated administrator username: %v", admin.ToString()))

	return nil
}

func (provider *Provider) UpdateAdminPassword(backend *utils.Backend, admin *models.Administrator, password string) error {
	// Validate password
	if len(password) == 0 {
		return fmt.Errorf("password cannot be empty")
	}

	// Verify admin exists
	if _, err := provider.GetAdminById(admin.Id); err != nil {
		return err
	}

	// Update admin
	update := database.UpdateAdminPasswordParams{
		ID:       admin.Id,
		Password: password,
	}
	dbadmin, err := provider.Queries.UpdateAdminPassword(context.Background(), update)
	if err != nil {
		backend.Log <- utils.NewLog(utils.Warn, fmt.Sprintf("Unable to update administrator <%s> password: %v", admin.Id, err.Error()))
		return err
	}
	a, err := ConvertDbToAdmin(dbadmin)
	if err != nil {
		return err
	}
	*admin = a
	backend.Log <- utils.NewLog(utils.Info, fmt.Sprintf("Updated administrator password: %v", admin.ToString()))

	return nil
}

func (provider *Provider) UpdateAdminRole(backend *utils.Backend, admin *models.Administrator, role utils.Role) error {
	// Validate the role
	if role == utils.InvalidRole {
		return fmt.Errorf("invalid role passed")
	}

	// Verify admin exists
	if _, err := provider.GetAdminById(admin.Id); err != nil {
		return err
	}

	// Update admin
	update := database.UpdateAdminRoleParams{
		ID:   admin.Id,
		Role: int32(role),
	}
	dbadmin, err := provider.Queries.UpdateAdminRole(context.Background(), update)
	if err != nil {
		backend.Log <- utils.NewLog(utils.Warn, fmt.Sprintf("Unable to update administrator <%s> role: %v", admin.Id, err.Error()))
		return err
	}
	a, err := ConvertDbToAdmin(dbadmin)
	if err != nil {
		return err
	}
	*admin = a
	backend.Log <- utils.NewLog(utils.Info, fmt.Sprintf("Updated administrator role: %v", admin.ToString()))

	return nil
}
