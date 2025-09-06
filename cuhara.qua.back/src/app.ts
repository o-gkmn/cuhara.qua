import express from 'express';
import { CQRSRegistry } from './core/cqrs/registry';
import userRoutes from './modules/users/user.routes';
import { errorHandler, notFoundHandler } from './middleware/error-handler';

const app = express();

// Initialize CQRS Registry
CQRSRegistry.getInstance();

// Middleware
app.use(express.json());
app.use(express.urlencoded({extended: true}));

// Routes
app.use('/api/users', userRoutes);

// Health check endpoint
app.get('/health', (req, res) => {
  res.json({ status: 'OK', timestamp: new Date().toISOString() });
});

// Error handling middleware (must be last)
app.use(notFoundHandler);
app.use(errorHandler);

export default app;