import { OpenAPI } from '../api';

// Token management utilities
export const setAuthToken = (token: string) => {
    localStorage.setItem('authToken', token);
    // Update OpenAPI config with the token
    OpenAPI.TOKEN = token;
};

export const getAuthToken = (): string | null => {
    return localStorage.getItem('authToken');
};

export const removeAuthToken = () => {
    localStorage.removeItem('authToken');
    OpenAPI.TOKEN = undefined;
};

export const isAuthenticated = (): boolean => {
    const token = getAuthToken();
    return token !== null && token !== '';
};

// Initialize token from localStorage on app start
export const initializeAuth = () => {
    const token = getAuthToken();
    if (token) {
        OpenAPI.TOKEN = token;
    }
};
