<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Register - Inventory Wizard</title>
    <!-- Bootstrap CSS -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
</head>
<body class="bg-light">
    <div class="container">
        <!-- App Branding -->
        <div class="text-center mt-5">
            <h1 class="display-4">Inventory Wizard</h1>
            <p class="lead">Create an account and manage your inventory with ease!</p>
        </div>

        <!-- Registration Form -->
        <div class="d-flex justify-content-center align-items-center mt-4">
            <div class="card shadow-sm" style="width: 400px;">
                <div class="card-body">
                    <h3 class="text-center mb-4">Register</h3>
                    <!-- Error Message -->
                    {{if .Error}}
                    <div class="alert alert-danger text-center" role="alert">
                        {{.Error}}
                    </div>
                    {{end}}
                    <!-- Form -->
                    <form action="/register" method="post">
                        <div class="mb-3">
                            <label for="username" class="form-label">Username</label>
                            <input type="text" class="form-control" id="username" name="username" placeholder="Enter a username" required>
                        </div>
                        <div class="mb-3">
                            <label for="email" class="form-label">Email</label>
                            <input type="email" class="form-control" id="email" name="email" placeholder="Enter your email" required>
                        </div>
                        <div class="mb-3">
                            <label for="password" class="form-label">Password</label>
                            <input type="password" class="form-control" id="password" name="password" placeholder="Create a password" required>
                        </div>
                        <div class="mb-3">
                            <label for="role" class="form-label">Role</label>
                            <select class="form-control" id="role" name="role" required>
                                <option value="user" selected>User</option>
                                <option value="admin">Admin</option>
                            </select>
                        </div>
                        <div class="d-grid">
                            <button type="submit" class="btn btn-primary">Register</button>
                        </div>
                    </form>
                    <div class="text-center mt-3">
                        <a href="/" class="btn btn-link">Already have an account? Login</a>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <!-- Footer -->
    <footer class="text-center mt-5">
        <p class="text-muted">&copy; 2024 Inventory Wizard. All Rights Reserved.</p>
    </footer>
    <!-- Bootstrap JS Bundle -->
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"></script>
</body>
</html>
