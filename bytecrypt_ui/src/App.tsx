import "./App.css";

import Header from "./components/Header";
import Content from "./components/Content";
import Footer from "./components/Footer";

export default function App() {

    return (
        <div className="App flex flex-col justify-between text-3xl text-white h-screen">
            <div className="mx-10 my-4 header flex justify-between">
                <Header />
            </div>
            <hr />

            <div className="content mx-10 self-start">
                <Content />
            </div>

            <div className="mt-auto">
                <hr />
                <div className="footer my-4 mx-10 flex justify-between">
                    <Footer />
                </div>
            </div>
        </div>
    );
}

