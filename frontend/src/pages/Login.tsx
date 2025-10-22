import { useState } from "react";
import { useNavigate } from "react-router-dom";
import MainLayout from "../components/layouts/MainLayout";
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
        <MainLayout>
            <div className="flex flex-col items-center justify-center min-h-full px-4 py-8">
                <div className="w-full max-w-md">
                    <div className="bg-white rounded-lg shadow-lg p-8">
                        <div className="text-center mb-8">
                            <h1 className="text-3xl font-bold text-gray-900 mb-2">Giriş Yap</h1>
                            <p className="text-gray-600">Hesabınıza giriş yapın</p>
                        </div>

                        {error && (
                            <div className="mb-4 p-3 bg-red-100 border border-red-400 text-red-700 rounded">
                                {error}
                            </div>
                        )}

                        <form onSubmit={handleSubmit} className="space-y-6">
                            <div>
                                <label htmlFor="email" className="block text-sm font-medium text-gray-700 mb-2">
                                    Email
                                </label>
                                <input
                                    type="email"
                                    id="email"
                                    name="email"
                                    value={formData.email}
                                    onChange={handleInputChange}
                                    placeholder="ornek@email.com"
                                    className="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                                    disabled={loading}
                                />
                            </div>

                            <div>
                                <label htmlFor="password" className="block text-sm font-medium text-gray-700 mb-2">
                                    Şifre
                                </label>
                                <input
                                    type="password"
                                    id="password"
                                    name="password"
                                    value={formData.password}
                                    onChange={handleInputChange}
                                    placeholder="Şifrenizi giriniz"
                                    className="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                                    disabled={loading}
                                />
                            </div>

                            <button
                                type="submit"
                                disabled={loading}
                                className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed"
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
                            <p className="text-sm text-gray-600">
                                Hesabınız yok mu?{" "}
                                <button 
                                    onClick={() => navigate("/register")}
                                    className="font-medium text-blue-600 hover:text-blue-500"
                                >
                                    Kayıt olun
                                </button>
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </MainLayout>
    )
}

export default Login;