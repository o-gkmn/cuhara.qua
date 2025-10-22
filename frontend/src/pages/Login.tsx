import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { AuthService } from "../api";
import type { loginRequest } from "../api";
import { setAuthToken } from "../utils/auth";

function Login() {
    const navigate = useNavigate();
    const [formData, setFormData] = useState<loginRequest>({
        email: "",
        password: ""
    });
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setFormData(prev => ({
            ...prev,
            [name]: value
        }));
        // Clear error when user starts typing
        if (error) setError(null);
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        
        // Basic validation
        if (!formData.email || !formData.password) {
            setError("Email ve şifre alanları zorunludur");
            return;
        }

        if (!formData.email.includes("@")) {
            setError("Geçerli bir email adresi giriniz");
            return;
        }

        try {
            setLoading(true);
            setError(null);
            
            const response = await AuthService.postApiV1AuthLogin(formData);
            
            if (response.token) {
                // Store token and update API config
                setAuthToken(response.token);
                
                // Redirect to home page
                navigate("/");
            } else {
                setError("Giriş başarısız. Lütfen bilgilerinizi kontrol ediniz.");
            }
        } catch (err) {
            console.error("Login error:", err);
            setError("Giriş başarısız. Email veya şifre hatalı.");
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="flex items-center justify-center min-h-screen w-full ">
            <div className="w-full max-w-md">
                <div className="bg-slate-800/60 border border-slate-700 rounded-lg shadow-xl p-8 backdrop-blur">
                    <div className="text-center mb-8">
                        <h1 className="text-3xl font-bold text-slate-100 mb-2">Giriş Yap</h1>
                        <p className="text-slate-300">Hesabınıza giriş yapın</p>
                    </div>

                    {error && (
                        <div className="mb-4 p-3 bg-red-900/50 border border-red-500 text-red-200 rounded">
                            {error}
                        </div>
                    )}

                    <form onSubmit={handleSubmit} className="space-y-6">
                        <div>
                            <label htmlFor="email" className="block text-sm font-medium text-slate-200 mb-2">
                                Email
                            </label>
                            <input
                                type="email"
                                id="email"
                                name="email"
                                value={formData.email}
                                onChange={handleInputChange}
                                placeholder="ornek@email.com"
                                className="w-full px-3 py-2 bg-slate-700 border border-slate-600 rounded-md shadow-sm text-slate-100 placeholder-slate-400 focus:outline-none focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500"
                                disabled={loading}
                            />
                        </div>

                        <div>
                            <label htmlFor="password" className="block text-sm font-medium text-slate-200 mb-2">
                                Şifre
                            </label>
                            <input
                                type="password"
                                id="password"
                                name="password"
                                value={formData.password}
                                onChange={handleInputChange}
                                placeholder="Şifrenizi giriniz"
                                className="w-full px-3 py-2 bg-slate-700 border border-slate-600 rounded-md shadow-sm text-slate-100 placeholder-slate-400 focus:outline-none focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500"
                                disabled={loading}
                            />
                        </div>

                        <button
                            type="submit"
                            disabled={loading}
                            className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-emerald-950 bg-emerald-500 hover:bg-emerald-400 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-emerald-500 disabled:opacity-50 disabled:cursor-not-allowed"
                        >
                            {loading ? (
                                <div className="flex items-center">
                                    <svg className="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                                        <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                                        <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                                    </svg>
                                    Giriş yapılıyor...
                                </div>
                            ) : (
                                "Giriş Yap"
                            )}
                        </button>
                    </form>

                    <div className="mt-6 text-center">
                        <p className="text-sm text-slate-300">
                            Hesabınız yok mu?{" "}
                            <button 
                                onClick={() => navigate("/register")}
                                className="font-medium text-emerald-400 hover:text-emerald-300"
                            >
                                Kayıt olun
                            </button>
                        </p>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Login;