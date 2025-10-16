import { BrowserRouter, Route, Routes } from "react-router-dom";
import Home from "../pages/Home";
import TablePage from "../pages/TablePage";

export default function AppRoutes() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/table" element={<TablePage />} />
            </Routes>
        </BrowserRouter>
    )
}