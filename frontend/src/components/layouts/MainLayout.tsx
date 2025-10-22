
import Navbar from "./Navbar";

export default function MainLayout({ children }: { children: React.ReactNode }) {
    return (
        <div className="min-h-screen flex flex-col bg-gradient-to-br from-slate-900 via-slate-900/95 to-slate-900/90">
            <Navbar />
            <main className="flex-1">
                {children}
            </main>
        </div>
    )
}