<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Warehouses - Inventory Wizard</title>
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
        <h1 class="text-center mb-4">Manage Warehouses</h1>
        <!-- Add Warehouse Button -->
        <div class="text-center mb-3">
            <button class="btn btn-success" data-bs-toggle="modal" data-bs-target="#addWarehouseModal">Add New Warehouse</button>
        </div>
        <!-- Warehouses Table -->
        <div class="table-responsive">
            <table class="table table-bordered table-striped">
                <thead class="table-dark">
                    <tr>
                        <th>ID</th>
                        <th>Name</th>
                        <th>Location</th>
                        <th>Manager</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {{range .Warehouses}}
                    <tr>
                        <td>{{.ID}}</td>
                        <td>{{.Name}}</td>
                        <td>{{.Location}}</td>
                        <td>{{if .Manager}}{{.Manager.Username}}{{else}}N/A{{end}}</td>
                        <td>
                            <button class="btn btn-primary btn-sm" data-bs-toggle="modal" data-bs-target="#editWarehouseModal" onclick="loadWarehouse({{.ID}}, '{{.Name}}', '{{.Location}}')">Edit</button>
                            <form method="post" action="/warehouse/{{.ID}}" class="d-inline">
                                <input type="hidden" name="_method" value="delete">
                                <button type="submit" class="btn btn-danger btn-sm">Delete</button>
                            </form>
                        </td>
                    </tr>
                    {{else}}
                    <tr>
                        <td colspan="5" class="text-center">No warehouses found.</td>
                    </tr>
                    {{end}}
                </tbody>
            </table>
        </div>
    </div>
    <!-- Add Warehouse Modal -->
    <div class="modal fade" id="addWarehouseModal" tabindex="-1" aria-labelledby="addWarehouseModalLabel" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title" id="addWarehouseModalLabel">Add New Warehouse</h5>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <form action="/warehouse" method="post">
                    <div class="modal-body">
                        <div class="mb-3">
                            <label for="name" class="form-label">Warehouse Name</label>
                            <input type="text" class="form-control" id="name" name="name" required>
                        </div>
                        <div class="mb-3">
                            <label for="location" class="form-label">Location</label>
                            <input type="text" class="form-control" id="location" name="location" required>
                        </div>
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Cancel</button>
                        <button type="submit" class="btn btn-success">Add Warehouse</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
    <!-- Edit Warehouse Modal -->
    <div class="modal fade" id="editWarehouseModal" tabindex="-1" aria-labelledby="editWarehouseModalLabel" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title" id="editWarehouseModalLabel">Edit Warehouse</h5>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <form id="editWarehouseForm" method="post">
                    <input type="hidden" name="_method" value="put">
                    <div class="modal-body">
                        <input type="hidden" id="editWarehouseID" name="id">
                        <div class="mb-3">
                            <label for="editName" class="form-label">Warehouse Name</label>
                            <input type="text" class="form-control" id="editName" name="name" required>
                        </div>
                        <div class="mb-3">
                            <label for="editLocation" class="form-label">Location</label>
                            <input type="text" class="form-control" id="editLocation" name="location" required>
                        </div>
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Cancel</button>
                        <button type="submit" class="btn btn-primary">Save Changes</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
    <!-- Footer -->
    <footer class="text-center mt-5">
        <p class="text-muted">&copy; 2024 Inventory Wizard. All Rights Reserved.</p>
    </footer>
    <!-- Bootstrap JS Bundle -->
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"></script>
    <script>
        function loadWarehouse(id, name, location) {
            const form = document.getElementById('editWarehouseForm');
            form.action = `/warehouse/${id}`;
            document.getElementById('editWarehouseID').value = id;
            document.getElementById('editName').value = name;
            document.getElementById('editLocation').value = location;
        }
    </script>
</body>
</html>