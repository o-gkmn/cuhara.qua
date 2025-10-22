import { Link, NavLink, useNavigate } from "react-router-dom";
import { removeAuthToken, isAuthenticated } from "../../utils/auth";
import { useState } from "react";

export default function Navbar() {
    const navigate = useNavigate();
    const authenticated = isAuthenticated();
    const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);

    const handleLogout = () => {
        removeAuthToken();
        navigate("/login");
        setIsMobileMenuOpen(false);
    };

    const toggleMobileMenu = () => {
        setIsMobileMenuOpen(!isMobileMenuOpen);
    };

    const closeMobileMenu = () => {
        setIsMobileMenuOpen(false);
    };

    return (
        <nav className="fixed top-0 left-0 right-0 z-50 bg-gradient-to-br from-slate-900 via-slate-900/95 to-slate-900/90 supports-[backdrop-filter]:bg-slate-900/60 backdrop-blur border-b border-slate-800/80">
            <div className="container mx-auto px-4">
                <div className="flex h-16 items-center justify-between">
                    {/* Logo */}
                    <Link to="/" className="inline-flex items-center gap-2" onClick={closeMobileMenu}>
                        <div className="w-8 h-8 bg-gradient-to-r from-emerald-400 to-cyan-400 rounded-lg flex items-center justify-center">
                            <span className="text-slate-900 font-bold text-sm">C</span>
                        </div>
                        <span className="text-xl font-extrabold bg-clip-text text-transparent bg-gradient-to-r from-emerald-400 to-cyan-400">
                            cuhara.qua
                        </span>
                    </Link>

                    {/* Desktop Navigation */}
                    <div className="hidden md:flex items-center">
                        <div className="rounded-xl bg-slate-800/60 border border-slate-700/70 shadow-xl backdrop-blur px-1.5 py-1 flex gap-1">
                            <NavLink
                                to="/"
                                end
                                className={({ isActive }) =>
                                    `px-4 py-2 rounded-lg text-sm font-medium transition-all duration-200
                                        ${isActive
                                        ? "bg-emerald-500 text-emerald-950 shadow-lg"
                                        : "text-slate-200 hover:text-white hover:bg-slate-700/70"}`
                                }
                            >
                                Home
                            </NavLink>
                            <NavLink
                                to="/table"
                                end
                                className={({ isActive }) =>
                                    `px-4 py-2 rounded-lg text-sm font-medium transition-all duration-200
                                        ${isActive
                                        ? "bg-emerald-500 text-emerald-950 shadow-lg"
                                        : "text-slate-200 hover:text-white hover:bg-slate-700/70"}`
                                }
                            >
                                Table
                            </NavLink>
                            {authenticated ? (
                                <button
                                    onClick={handleLogout}
                                    className="px-4 py-2 rounded-lg text-sm font-medium transition-all duration-200 text-slate-200 hover:text-white hover:bg-slate-700/70"
                                >
                                    Logout
                                </button>
                            ) : (
                                <NavLink
                                    to="/login"
                                    end
                                    className={({ isActive }) =>
                                        `px-4 py-2 rounded-lg text-sm font-medium transition-all duration-200
                                            ${isActive
                                            ? "bg-emerald-500 text-emerald-950 shadow-lg"
                                            : "text-slate-200 hover:text-white hover:bg-slate-700/70"}`
                                    }
                                >
                                    Login
                                </NavLink>
                            )}
                        </div>
                    </div>

                    {/* Mobile menu button */}
                    <button
                        onClick={toggleMobileMenu}
                        className="md:hidden p-2 rounded-lg text-slate-200 hover:text-white hover:bg-slate-700/70 transition-colors"
                        aria-label="Toggle mobile menu"
                    >
                        <svg
                            className={`w-6 h-6 transition-transform duration-200 ${isMobileMenuOpen ? 'rotate-90' : ''}`}
                            fill="none"
                            stroke="currentColor"
                            viewBox="0 0 24 24"
                        >
                            {isMobileMenuOpen ? (
                                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
                            ) : (
                                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 6h16M4 12h16M4 18h16" />
                            )}
                        </svg>
                    </button>
                </div>

                {/* Mobile Navigation */}
                {isMobileMenuOpen && (
                    <div className="md:hidden border-t border-slate-700/70 bg-slate-800/90 backdrop-blur">
                        <div className="px-2 py-4 space-y-2">
                            <NavLink
                                to="/"
                                end
                                onClick={closeMobileMenu}
                                className={({ isActive }) =>
                                    `block px-4 py-3 rounded-lg text-sm font-medium transition-all duration-200
                                        ${isActive
                                        ? "bg-emerald-500 text-emerald-950 shadow-lg"
                                        : "text-slate-200 hover:text-white hover:bg-slate-700/70"}`
                                }
                            >
                                Home
                            </NavLink>
                            <NavLink
                                to="/table"
                                end
                                onClick={closeMobileMenu}
                                className={({ isActive }) =>
                                    `block px-4 py-3 rounded-lg text-sm font-medium transition-all duration-200
                                        ${isActive
                                        ? "bg-emerald-500 text-emerald-950 shadow-lg"
                                        : "text-slate-200 hover:text-white hover:bg-slate-700/70"}`
                                }
                            >
                                Table
                            </NavLink>
                            {authenticated ? (
                                <button
                                    onClick={handleLogout}
                                    className="block w-full text-left px-4 py-3 rounded-lg text-sm font-medium transition-all duration-200 text-slate-200 hover:text-white hover:bg-slate-700/70"
                                >
                                    Logout
                                </button>
                            ) : (
                                <NavLink
                                    to="/login"
                                    end
                                    onClick={closeMobileMenu}
                                    className={({ isActive }) =>
                                        `block px-4 py-3 rounded-lg text-sm font-medium transition-all duration-200
                                            ${isActive
                                            ? "bg-emerald-500 text-emerald-950 shadow-lg"
                                            : "text-slate-200 hover:text-white hover:bg-slate-700/70"}`
                                    }
                                >
                                    Login
                                </NavLink>
                            )}
                        </div>
                    </div>
                )}
            </div>
        </nav>
    )
}   