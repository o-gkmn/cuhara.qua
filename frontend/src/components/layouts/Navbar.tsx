import { Link, NavLink, useNavigate } from "react-router-dom";
import { removeAuthToken, isAuthenticated } from "../../utils/auth";

export default function Navbar() {
    const navigate = useNavigate();
    const authenticated = isAuthenticated();

    const handleLogout = () => {
        removeAuthToken();
        navigate("/login");
    };

    return (
        <nav className="fixed top-0 left-0 right-0 z-50 bg-gradient-to-br from-slate-900 via-slate-900/95 to-slate-900/90 supports-[backdrop-filter]:bg-slate-900/60 backdrop-blur border-b border-slate-800/80">
            <div className="container mx-auto px-4">
                <div className="flex h-14 items-center justify-between">
                    <Link to="/" className="inline-flex items-center gap-2">
                        <span className="text-lg font-extrabold bg-clip-text text-transparent bg-gradient-to-r from-emerald-400 to-cyan-400">
                            cuhara.qua
                        </span>
                    </Link>

                    <div className="flex items-center">
                        <div className="rounded-xl bg-slate-800/60 border border-slate-700/70 shadow-xl backdrop-blur px-1.5 py-1 flex gap-4" >
                            <NavLink
                                to="/"
                                end
                                className={({ isActive }) =>
                                    `px-3 py-1.5 rounded-lg text-sm font-medium transition
                                        ${isActive
                                        ? "bg-emerald-500 text-emerald-950 shadow"
                                        : "text-slate-200 hover:text-white hover:bg-slate-700/70"}`
                                }
                            >
                                Home
                            </NavLink>
                            <NavLink
                                to="/table"
                                end
                                className={({ isActive }) =>
                                    `px-3 py-1.5 rounded-lg text-sm font-medium transition
                                        ${isActive
                                        ? "bg-emerald-500 text-emerald-950 shadow"
                                        : "text-slate-200 hover:text-white hover:bg-slate-700/70"}`
                                }
                            >
                                Table
                            </NavLink>
                            {authenticated ? (
                                <button
                                    onClick={handleLogout}
                                    className="px-3 py-1.5 rounded-lg text-sm font-medium transition text-slate-200 hover:text-white hover:bg-slate-700/70"
                                >
                                    Logout
                                </button>
                            ) : (
                                <NavLink
                                    to="/login"
                                    end
                                    className={({ isActive }) =>
                                        `px-3 py-1.5 rounded-lg text-sm font-medium transition
                                            ${isActive
                                            ? "bg-emerald-500 text-emerald-950 shadow"
                                            : "text-slate-200 hover:text-white hover:bg-slate-700/70"}`
                                    }
                                >
                                    Login
                                </NavLink>
                            )}
                        </div>
                    </div>
                </div>
            </div>
        </nav>
    )
}   