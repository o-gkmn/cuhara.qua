import { BrowserRouter, Route, Routes } from "react-router-dom";
import Home from "../pages/Home";
import TablePage from "../pages/TablePage";
import Login from "../pages/Login";

export default function AppRoutes() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/table" element={<TablePage />} />
                <Route path="/login" element={<Login />} />
            </Routes>
        </BrowserRouter>
    )
}