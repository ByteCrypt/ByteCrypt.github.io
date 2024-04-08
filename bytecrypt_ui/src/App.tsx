import "./App.css";

import Header from "./components/Header";
import Content from "./components/Content";
import Footer from "./components/Footer";

export default function App() {

    return (
        <>
            <div className="text-white fixed top-0 w-full flex flex-col">
                <div className="my-4 header flex sm:flex-row flex-col sm:justify-between justify-center mx-auto sm:mx-10">
                    <Header />
                </div>
                <hr />
            </div>
            <div className="App flex flex-col justify-between text-3xl text-white h-screen pt-24">
                <div className="content mx-auto self-start my-10 flex flex-col md:flex-row justify-center">
                    <Content />
                </div>

                <div className="mt-auto">
                    <hr />
                    <div className="footer my-4 mx-10 flex sm:flex-row flex-col sm:justify-between">
                        <Footer />
                    </div>
                </div>
            </div>
        </>
    );
}

