<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Home - Inventory Wizard</title>
    <!-- Bootstrap CSS -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
</head>
<body class="bg-light">
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
        <div class="container">
            <a class="navbar-brand" href="/home">Inventory Wizard</a>
            <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarNav">
                <ul class="navbar-nav ms-auto">
                    <li class="nav-item">
                        <a class="nav-link" href="/logout">Logout</a>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
    <div class="container mt-5">
        <h1 class="text-center mb-4">Welcome to Inventory Wizard!</h1>
        <p class="text-center lead">Manage your app structures easily with the tools below:</p>
        <div class="row">
            <div class="col-md-6 mb-3">
                <div class="card shadow-sm">
                    <div class="card-body text-center">
                        <h5 class="card-title">Products</h5>
                        <p class="card-text">Manage all your products here.</p>
                        <a href="/products" class="btn btn-primary">Go to Products</a>
                    </div>
                </div>
            </div>
            <div class="col-md-6 mb-3">
                <div class="card shadow-sm">
                    <div class="card-body text-center">
                        <h5 class="card-title">Stocks</h5>
                        <p class="card-text">Manage stock levels and inventory.</p>
                        <a href="/stocks" class="btn btn-primary">Go to Stocks</a>
                    </div>
                </div>
            </div>
            <div class="col-md-6 mb-3">
                <div class="card shadow-sm">
                    <div class="card-body text-center">
                        <h5 class="card-title">Warehouses</h5>
                        <p class="card-text">Manage warehouse locations and details.</p>
                        <a href="/warehouses" class="btn btn-primary">Go to Warehouses</a>
                    </div>
                </div>
            </div>
            <div class="col-md-6 mb-3">
                <div class="card shadow-sm">
                    <div class="card-body text-center">
                        <h5 class="card-title">Users</h5>
                        <p class="card-text">Manage user accounts and roles.</p>
                        <a href="/users" class="btn btn-primary">Go to Users</a>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <footer class="text-center mt-5">
        <p class="text-muted">&copy; 2024 Inventory Wizard. All Rights Reserved.</p>
    </footer>
    <!-- Bootstrap JS Bundle -->
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"></script>
</body>
</html>
