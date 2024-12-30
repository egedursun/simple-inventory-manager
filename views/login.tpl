<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Inventory Wizard - Login</title>
    <!-- Bootstrap CSS -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
</head>
<body class="bg-light">
    <div class="container">
        <!-- App Branding -->
        <div class="text-center mt-5">
            <h1 class="display-4">Inventory Wizard</h1>
            <p class="lead">Manage your inventory seamlessly</p>
        </div>

        <!-- Login Form -->
        <div class="d-flex justify-content-center align-items-center mt-4">
            <div class="card shadow-sm" style="width: 400px;">
                <div class="card-body">
                    <h3 class="text-center mb-4">Login</h3>
                    <!-- Error Message -->
                    {{if .Error}}
                    <div class="alert alert-danger text-center" role="alert">
                        {{.Error}}
                    </div>
                    {{end}}
                    <!-- Form -->
                    <form action="/login" method="post">
                        <div class="mb-3">
                            <label for="username" class="form-label">Username</label>
                            <input type="text" class="form-control" id="username" name="username" placeholder="Enter your username" required>
                        </div>
                        <div class="mb-3">
                            <label for="password" class="form-label">Password</label>
                            <input type="password" class="form-control" id="password" name="password" placeholder="Enter your password" required>
                        </div>
                        <div class="d-grid">
                            <button type="submit" class="btn btn-primary">Login</button>
                        </div>
                    </form>
                    <!-- Register Link -->
                    <div class="text-center mt-3">
                        <p>Don't have an account? <a href="/register" class="btn btn-link">Register</a></p>
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