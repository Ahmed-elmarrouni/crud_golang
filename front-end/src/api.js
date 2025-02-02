import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080';

// Fetch all users
export const fetchUsers = async () => {
    try {
        const response = await axios.get(`${API_BASE_URL}/users`);
        return response.data;
    } catch (error) {
        console.error('err fetching users', error);
        throw error;
    }
}

// Create a user
export const addUser = async (user) => {
    try {
        const response = await axios.post(`${API_BASE_URL}/users`, user);
        return response.data;
    } catch (error) {
        console.error("Err creating user", error);
        throw error;
    }
};
