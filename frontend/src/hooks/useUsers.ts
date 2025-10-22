import { useState, useEffect } from 'react';
import { UsersService } from '../api';
import type { userResponse } from '../api';

interface UseUsersReturn {
  users: userResponse[];
  loading: boolean;
  error: string | null;
  refetch: () => void;
}

export const useUsers = (): UseUsersReturn => {
  const [users, setUsers] = useState<userResponse[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  const fetchUsers = async () => {
    try {
      setLoading(true);
      setError(null);
      const response = await UsersService.getApiV1Users();
      // API returns a single userResponse, but we need an array
      // If the API actually returns an array, we can directly set it
      // For now, let's assume it returns a single user and wrap it in an array
      setUsers(Array.isArray(response) ? response : [response]);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to fetch users');
      console.error('Error fetching users:', err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchUsers();
  }, []);

  return {
    users,
    loading,
    error,
    refetch: fetchUsers,
  };
};
