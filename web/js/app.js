// API Base URL
const API_URL = 'http://localhost:8080';

// State
let currentUser = null;
let authToken = localStorage.getItem('token') || null;

// DOM Elements
const authPage = document.getElementById('auth-page');
const dashboardPage = document.getElementById('dashboard-page');
const loginForm = document.getElementById('login-form');
const registerForm = document.getElementById('register-form');
const taskForm = document.getElementById('task-form');
const tasksContainer = document.getElementById('tasks-container');
const authMessage = document.getElementById('auth-message');
const dashboardMessage = document.getElementById('dashboard-message');

// Initialize
document.addEventListener('DOMContentLoaded', () => {
    setupEventListeners();
    if (authToken) {
        showDashboard();
        loadTasks();
    }
});

// Event Listeners
function setupEventListeners() {
    // Tabs
    document.querySelectorAll('.tab').forEach(tab => {
        tab.addEventListener('click', () => {
            const tabName = tab.dataset.tab;
            document.querySelectorAll('.tab').forEach(t => t.classList.remove('active'));
            document.querySelectorAll('.auth-form').forEach(f => f.classList.remove('active'));
            tab.classList.add('active');
            document.getElementById(`${tabName}-form`).classList.add('active');
        });
    });

    // Forms
    loginForm.addEventListener('submit', handleLogin);
    registerForm.addEventListener('submit', handleRegister);
    taskForm.addEventListener('submit', handleCreateTask);

    // Buttons
    document.getElementById('logout-btn').addEventListener('click', handleLogout);
    document.getElementById('add-task-btn').addEventListener('click', showTaskForm);
    document.getElementById('cancel-task-btn').addEventListener('click', hideTaskForm);
}

// Auth Functions
async function handleRegister(e) {
    e.preventDefault();
    const name = document.getElementById('register-name').value;
    const email = document.getElementById('register-email').value;
    const password = document.getElementById('register-password').value;

    try {
        const response = await fetch(`${API_URL}/register`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ name, email, password })
        });

        const data = await response.json();
        
        if (response.ok) {
            showMessage(authMessage, 'Registration successful! Please login.', 'success');
            registerForm.reset();
            // Switch to login tab
            document.querySelector('[data-tab="login"]').click();
        } else {
            showMessage(authMessage, data.message || 'Registration failed', 'error');
        }
    } catch (error) {
        showMessage(authMessage, 'Network error. Please try again.', 'error');
    }
}

async function handleLogin(e) {
    e.preventDefault();
    const email = document.getElementById('login-email').value;
    const password = document.getElementById('login-password').value;

    try {
        const response = await fetch(`${API_URL}/login`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ email, password })
        });

        const data = await response.json();
        
        if (response.ok) {
            authToken = data.token;
            localStorage.setItem('token', authToken);
            currentUser = { email };
            showDashboard();
            loadTasks();
        } else {
            showMessage(authMessage, 'Invalid credentials', 'error');
        }
    } catch (error) {
        showMessage(authMessage, 'Network error. Please try again.', 'error');
    }
}

function handleLogout() {
    authToken = null;
    currentUser = null;
    localStorage.removeItem('token');
    showAuth();
    loginForm.reset();
    registerForm.reset();
}

// Task Functions
async function loadTasks() {
    tasksContainer.innerHTML = '<div class="loading">Loading tasks...</div>';

    try {
        const response = await fetch(`${API_URL}/tasks`, {
            headers: { 'Authorization': `Bearer ${authToken}` }
        });

        if (response.ok) {
            const tasks = await response.json();
            displayTasks(tasks || []);
        } else {
            tasksContainer.innerHTML = '<div class="loading">Failed to load tasks</div>';
        }
    } catch (error) {
        tasksContainer.innerHTML = '<div class="loading">Network error</div>';
    }
}

function displayTasks(tasks) {
    if (tasks.length === 0) {
        tasksContainer.innerHTML = '<div class="loading">No tasks yet. Create one!</div>';
        return;
    }

    tasksContainer.innerHTML = tasks.map(task => `
        <div class="task-card">
            <div class="task-header">
                <div class="task-title">${escapeHtml(task.title)}</div>
                <div class="task-actions">
                    <button class="btn btn-secondary" onclick="updateTaskStatus(${task.id}, '${getNextStatus(task.status)}')">
                        ${getStatusButtonText(task.status)}
                    </button>
                    <button class="btn btn-secondary" onclick="deleteTask(${task.id})">Delete</button>
                </div>
            </div>
            ${task.description ? `<div class="task-description">${escapeHtml(task.description)}</div>` : ''}
            <div class="task-footer">
                <span class="task-status status-${task.status}">${formatStatus(task.status)}</span>
                <span class="task-date">${formatDate(task.created_at)}</span>
            </div>
        </div>
    `).join('');
}

async function handleCreateTask(e) {
    e.preventDefault();
    const title = document.getElementById('task-title').value;
    const description = document.getElementById('task-description').value;

    try {
        const response = await fetch(`${API_URL}/tasks`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${authToken}`
            },
            body: JSON.stringify({ title, description })
        });

        if (response.ok) {
            showMessage(dashboardMessage, 'Task created successfully!', 'success');
            taskForm.reset();
            hideTaskForm();
            loadTasks();
        } else {
            showMessage(dashboardMessage, 'Failed to create task', 'error');
        }
    } catch (error) {
        showMessage(dashboardMessage, 'Network error', 'error');
    }
}

async function updateTaskStatus(taskId, newStatus) {
    try {
        const response = await fetch(`${API_URL}/tasks/update?id=${taskId}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${authToken}`
            },
            body: JSON.stringify({ status: newStatus })
        });

        if (response.ok) {
            loadTasks();
        } else {
            showMessage(dashboardMessage, 'Failed to update task', 'error');
        }
    } catch (error) {
        showMessage(dashboardMessage, 'Network error', 'error');
    }
}

async function deleteTask(taskId) {
    if (!confirm('Are you sure you want to delete this task?')) return;

    try {
        const response = await fetch(`${API_URL}/tasks/delete?id=${taskId}`, {
            method: 'DELETE',
            headers: { 'Authorization': `Bearer ${authToken}` }
        });

        if (response.ok) {
            showMessage(dashboardMessage, 'Task deleted', 'success');
            loadTasks();
        } else {
            showMessage(dashboardMessage, 'Failed to delete task', 'error');
        }
    } catch (error) {
        showMessage(dashboardMessage, 'Network error', 'error');
    }
}

// UI Functions
function showAuth() {
    authPage.classList.add('active');
    dashboardPage.classList.remove('active');
}

function showDashboard() {
    authPage.classList.remove('active');
    dashboardPage.classList.add('active');
    document.getElementById('user-name').textContent = currentUser?.email || 'User';
}

function showTaskForm() {
    document.getElementById('task-form-container').style.display = 'block';
}

function hideTaskForm() {
    document.getElementById('task-form-container').style.display = 'none';
    taskForm.reset();
}

function showMessage(element, text, type) {
    element.textContent = text;
    element.className = `message ${type}`;
    setTimeout(() => {
        element.className = 'message';
    }, 3000);
}

// Utility Functions
function escapeHtml(text) {
    const div = document.createElement('div');
    div.textContent = text;
    return div.innerHTML;
}

function formatStatus(status) {
    return status.replace('_', ' ').toUpperCase();
}

function formatDate(dateString) {
    return new Date(dateString).toLocaleDateString('tr-TR');
}

function getNextStatus(currentStatus) {
    const statusFlow = {
        'pending': 'in_progress',
        'in_progress': 'done',
        'done': 'pending'
    };
    return statusFlow[currentStatus] || 'pending';
}

function getStatusButtonText(currentStatus) {
    const buttonText = {
        'pending': 'Start',
        'in_progress': 'Complete',
        'done': 'Reset'
    };
    return buttonText[currentStatus] || 'Update';
}
