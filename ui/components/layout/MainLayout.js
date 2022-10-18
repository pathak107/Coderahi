import Footer from "./footer";
import Header from "./header";

const MainLayout = ({ children }) => {
    return (
        <>
            <div className="mx-28">
                <Header />
                {children}
                
            </div>
            <Footer/>
        </>
    );
}

export default MainLayout;