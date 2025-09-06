import { Router } from "express";
import { UserController } from "./user.controller";
import { validateCreateUser, validateId } from "../../middleware/validation";

const router = Router();
const userController = new UserController();

// User routes
router.post("/", validateCreateUser, userController.createUser.bind(userController));
router.get("/:id", validateId, userController.getUser.bind(userController));
export default router;