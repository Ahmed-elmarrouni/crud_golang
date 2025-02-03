import React, { useEffect, useState } from 'react';
import {
  Container,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  TextField,
  Typography,
  Box,
  // Image,
} from '@mui/material';

import { fetchUsers, addUser, updateUser, deleteUser } from './api';

const App = () => {
  const [users, setUsers] = useState([]);
  const [open, setOpen] = useState(false);
  // const [newUser, setNewUser] = useState({ name: '', email: '' });
  // const [editingUser, setEditingUser] = useState(null);
  const [newUser, setNewUser] = useState({ name: '', email: '', image_url: '' });
  const [editingUser, setEditingUser] = useState(null);

  const [searchTerm, setSearchTerm] = useState('');
  // const [deleteTerm, setDeleteTern] = useState('')

  useEffect(() => {
    loadUsers();
  }, []);

  const loadUsers = async () => {
    try {
      const data = await fetchUsers();
      setUsers(data);
    } catch (error) {
      console.error('Error loading users:', error);
    }
  };


  const handleEditOpen = (user) => {
    setEditingUser(user);
    setOpen(true);
  };

  const handleEditClose = () => {
    setOpen(false);
    setEditingUser(null);
  };

  const handleAddUser = async () => {
    if (!newUser.name || !newUser.email) {
      alert("Both Name and Email are required!");
      return;
    }
    try {
      await addUser(newUser);
      setNewUser({ name: "", email: "", image_url: "" });
      setOpen(false);
      loadUsers();
    } catch (error) {
      console.error("Error adding user:", error);
    }
  };

  const handleEditSave = async () => {
    if (!editingUser.name || !editingUser.email) {
      alert("Both Name and Email are required!");
      return;
    }
    try {
      await updateUser(editingUser);
      setOpen(false);
      loadUsers();
    } catch (error) {
      console.error("Error updating user:", error);
    }
  };

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    if (editingUser) {
      setEditingUser((prev) => ({ ...prev, [name]: value }));
    } else {
      setNewUser((prev) => ({ ...prev, [name]: value }));
    }
  };

  const handleSearchChange = (event) => {
    setSearchTerm(event.target.value);
  };

  const filteredUsers = users.filter((user) =>
    user.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
    user.email.toLowerCase().includes(searchTerm.toLowerCase())
  );

  const handleDelete = async (userId) => {
    try {
      await deleteUser({ id: userId });
      await deleteUser({ id: userId });
      loadUsers();

    } catch (error) {
      console.error('Err deleting user:', error);

    }
  }

  const handleImageUpload = (event) => {
    const file = event.target.files[0];
    if (file) {
      const reader = new FileReader();
      reader.onloadend = () => {
        if (editingUser) {
          setEditingUser((prev) => ({ ...prev, image_url: reader.result }));
        } else {
          setNewUser((prev) => ({ ...prev, image_url: reader.result }));
        }
      };
      reader.readAsDataURL(file);
    }
  };

  return (
    <Container>
      <Typography variant="h4" align="center" gutterBottom>
        User Management
      </Typography>
      <div style={{ display: 'flex', justifyContent: 'space-between', marginBottom: 30, marginTop: 40 }}>
        <Box sx={{ width: 500, maxWidth: '100%' }}>
          <TextField
            fullWidth
            label="Search"
            value={searchTerm}
            onChange={handleSearchChange}
          />
        </Box>
        <Button variant="contained" color="primary" onClick={() => setOpen(true)} style={{ marginBottom: 20 }}>
          Add User
        </Button>
      </div>
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell><strong>ID</strong></TableCell>
              <TableCell><strong>Image</strong></TableCell>
              <TableCell><strong>Name</strong></TableCell>
              <TableCell><strong>Email</strong></TableCell>
              <TableCell><strong>Actions</strong></TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {filteredUsers.map((user) => (
              <TableRow key={user.id}>
                <TableCell>{user.id}</TableCell>
                <TableCell>
                  {user.image_url ? (
                    <img
                      src={user.image_url}
                      alt={user.image_url}
                      style={{ width: "3rem", height: "3rem", borderRadius: "50px" }}
                    />
                  ) : (
                    <Typography variant="body2" color="textSecondary">No Image</Typography>
                  )}
                </TableCell>
                <TableCell>{user.name}</TableCell>
                <TableCell>{user.email}</TableCell>
                <TableCell>
                  <Button variant="contained" color="primary" onClick={() => handleEditOpen(user)} style={{ marginRight: 10 }}>
                    Edit
                  </Button>
                  <Button variant="contained" color="secondary"
                    onClick={() => handleDelete(user.id)}
                  >
                    Delete
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
      <Dialog open={open} onClose={handleEditClose}>
        <DialogTitle>{editingUser ? 'Edit User' : 'Add User'}</DialogTitle>
        <DialogContent>
          <TextField margin="dense" label="Name" name="name" value={editingUser ? editingUser.name : newUser.name} onChange={handleInputChange} fullWidth />
          <TextField margin="dense" label="Email" name="email" value={editingUser ? editingUser.email : newUser.email} onChange={handleInputChange} fullWidth />
          <input
            type="file"
            accept="image/*"
            onChange={handleImageUpload}
            style={{ marginTop: 15 }}
          />

          {/* Preview Selected Image */}
          {((editingUser && editingUser.image_url) || newUser.image_url) && (
            <img
              src={editingUser ? editingUser.image_url : newUser.image_url}
              alt="Preview"
              style={{ width: "100px", height: "100px", marginTop: 10, borderRadius: "5px" }}
            />
          )}

        </DialogContent>
        <DialogActions>
          <Button onClick={handleEditClose} color="secondary">Cancel</Button>
          <Button onClick={editingUser ? handleEditSave : handleAddUser} color="primary">{editingUser ? 'Save' : 'Add'}</Button>
        </DialogActions>
      </Dialog>
    </Container>
  );
};

export default App;
