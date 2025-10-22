export default function Footer() {
    return (
        <footer className="fixed bottom-0 left-0 right-0 z-50 border-t border-slate-800/80 bg-gradient-to-tr from-slate-900 via-slate-900/95 to-slate-900/90 supports-[backdrop-filter]:bg-slate-900/60 backdrop-blur">
            <div className="container mx-auto px-4">
                <div className="h-14 flex items-center justify-between text-sm text-slate-400">
                    <span>© {new Date().getFullYear()} cuhara.qua</span>
                </div>
            </div>
        </footer>
    )
}