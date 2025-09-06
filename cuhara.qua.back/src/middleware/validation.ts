import { Request, Response, NextFunction } from 'express';

export const validateCreateUser = (req: Request, res: Response, next: NextFunction): void => {
  const { name, email, vscAccount, roleId, tenantId } = req.body;

  if (!name || !email || !vscAccount || !roleId || !tenantId) {
    res.status(400).json({
      success: false,
      error: 'Missing required fields: name, email, vscAccount, roleId, tenantId',
    });
    return;
  }

  if (typeof name !== 'string' || typeof email !== 'string' || typeof vscAccount !== 'string') {
    res.status(400).json({
      success: false,
      error: 'name, email, and vscAccount must be strings',
    });
    return;
  }

  if (typeof roleId !== 'number' || typeof tenantId !== 'number') {
    res.status(400).json({
      success: false,
      error: 'roleId and tenantId must be numbers',
    });
    return;
  }

  // Email validation
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  if (!emailRegex.test(email)) {
    res.status(400).json({
      success: false,
      error: 'Invalid email format',
    });
    return;
  }

  next();
};

export const validateCreatePost = (req: Request, res: Response, next: NextFunction): void => {
  const { creatorId, subtopicId, tenantId } = req.body;

  if (!creatorId || !subtopicId || !tenantId) {
    res.status(400).json({
      success: false,
      error: 'Missing required fields: creatorId, subtopicId, tenantId',
    });
    return;
  }

  if (typeof creatorId !== 'number' || typeof subtopicId !== 'number' || typeof tenantId !== 'number') {
    res.status(400).json({
      success: false,
      error: 'creatorId, subtopicId, and tenantId must be numbers',
    });
    return;
  }

  if (req.body.tagIds && !Array.isArray(req.body.tagIds)) {
    res.status(400).json({
      success: false,
      error: 'tagIds must be an array',
    });
    return;
  }

  next();
};

export const validateId = (req: Request, res: Response, next: NextFunction): void => {
  const { id } = req.params;
  const numId = parseInt(id);

  if (isNaN(numId) || numId <= 0) {
    res.status(400).json({
      success: false,
      error: 'Invalid ID parameter',
    });
    return;
  }

  next();
};