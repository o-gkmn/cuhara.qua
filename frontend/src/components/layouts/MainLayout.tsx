import Footer from "./Footer";
import Navbar from "./Navbar";

export default function MainLayout({ children }: { children: React.ReactNode }) {
    return (
        <div className="bg-gray-100">
            <Navbar/>
            <main className="pt-14 pb-14 min-h-[calc(100vh-7rem)]">{children}</main>
            <Footer/>
        </div>
    )
}