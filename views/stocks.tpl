<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Stocks - Inventory Wizard</title>
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
        <h1 class="text-center mb-4">Manage Stocks</h1>
        <!-- Add Stock Button -->
        <div class="text-center mb-3">
            <button class="btn btn-success" data-bs-toggle="modal" data-bs-target="#addStockModal">Add New Stock</button>
        </div>
        <!-- Stocks Table -->
        <div class="table-responsive">
            <table class="table table-bordered table-striped">
                <thead class="table-dark">
                    <tr>
                        <th>ID</th>
                        <th>Warehouse</th>
                        <th>Product</th>
                        <th>Quantity</th>
                        <th>Threshold</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {{range .Stocks}}
                    <tr>
                        <td>{{.ID}}</td>
                        <td>{{.Warehouse.Name}}</td>
                        <td>{{.Product.Name}}</td>
                        <td>{{.Quantity}}</td>
                        <td>{{.Threshold}}</td>
                        <td>
                            <button class="btn btn-primary btn-sm" data-bs-toggle="modal" data-bs-target="#editStockModal" onclick="loadStock({{.ID}}, '{{.Warehouse.ID}}', '{{.Product.ID}}', '{{.Quantity}}', '{{.Threshold}}')">Edit</button>
                            <button class="btn btn-danger btn-sm" onclick="deleteStock({{.ID}})">Delete</button>
                        </td>
                    </tr>
                    {{else}}
                    <tr>
                        <td colspan="6" class="text-center">No stocks found.</td>
                    </tr>
                    {{end}}
                </tbody>
            </table>
        </div>
    </div>
    <!-- Add Stock Modal -->
    <div class="modal fade" id="addStockModal" tabindex="-1" aria-labelledby="addStockModalLabel" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title" id="addStockModalLabel">Add New Stock</h5>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <form action="/stock" method="post">
                    <div class="modal-body">
                        <div class="mb-3">
                            <label for="warehouse" class="form-label">Warehouse</label>
                            <select id="warehouse" name="warehouse_id" class="form-control" required>
                                {{range .Warehouses}}
                                <option value="{{.ID}}">{{.Name}}</option>
                                {{end}}
                            </select>
                        </div>
                        <div class="mb-3">
                            <label for="product" class="form-label">Product</label>
                            <select id="product" name="product_id" class="form-control" required>
                                {{range .Products}}
                                <option value="{{.ID}}">{{.Name}}</option>
                                {{end}}
                            </select>
                        </div>
                        <div class="mb-3">
                            <label for="quantity" class="form-label">Quantity</label>
                            <input type="number" class="form-control" id="quantity" name="quantity" required>
                        </div>
                        <div class="mb-3">
                            <label for="threshold" class="form-label">Threshold</label>
                            <input type="number" class="form-control" id="threshold" name="threshold" required>
                        </div>
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Cancel</button>
                        <button type="submit" class="btn btn-success">Add Stock</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
    <!-- Edit Stock Modal -->
    <div class="modal fade" id="editStockModal" tabindex="-1" aria-labelledby="editStockModalLabel" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title" id="editStockModalLabel">Edit Stock</h5>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <form id="editStockForm" method="post">
                    <input type="hidden" name="_method" value="put">
                    <div class="modal-body">
                        <input type="hidden" id="editStockID" name="id">
                        <div class="mb-3">
                            <label for="editWarehouseID" class="form-label">Warehouse ID</label>
                            <input type="number" class="form-control" id="editWarehouseID" name="warehouse_id" required>
                        </div>
                        <div class="mb-3">
                            <label for="editProductID" class="form-label">Product ID</label>
                            <input type="number" class="form-control" id="editProductID" name="product_id" required>
                        </div>
                        <div class="mb-3">
                            <label for="editQuantity" class="form-label">Quantity</label>
                            <input type="number" class="form-control" id="editQuantity" name="quantity" required>
                        </div>
                        <div class="mb-3">
                            <label for="editThreshold" class="form-label">Threshold</label>
                            <input type="number" class="form-control" id="editThreshold" name="threshold" required>
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
        let currentStockId = null;

        function loadStock(id, warehouseID, productID, quantity, threshold) {
            const form = document.getElementById('editStockForm');
            form.action = `/stock/${id}`;
            document.getElementById('editWarehouseID').value = warehouseID;
            document.getElementById('editProductID').value = productID;
            document.getElementById('editQuantity').value = quantity;
            document.getElementById('editThreshold').value = threshold;
        }

        async function submitEditStock(event) {
            event.preventDefault();

            const warehouseId = document.getElementById('editWarehouse').value;
            const productId = document.getElementById('editProduct').value;
            const quantity = document.getElementById('editQuantity').value;
            const threshold = document.getElementById('editThreshold').value;

            const payload = { warehouse_id: warehouseId, product_id: productId, quantity, threshold };

            try {
                const response = await fetch(`/stock/${currentStockId}`, {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(payload),
                });

                if (!response.ok) {
                    const errorData = await response.json();
                    alert(`Failed to update stock: ${errorData.error || 'Unknown error'}`);
                    return;
                }

                alert("Stock updated successfully!");
                window.location.href = "/stocks";
            } catch (error) {
                alert("An error occurred while updating the stock.");
            }
        }

        async function deleteStock(id) {
            if (confirm("Are you sure you want to delete this stock?")) {
                try {
                    const response = await fetch(`/stock/${id}`, {
                        method: "DELETE",
                        headers: {
                            "Content-Type": "application/json",
                        },
                    });

                    if (!response.ok) {
                        const errorData = await response.json();
                        alert(`Failed to delete stock: ${errorData.error || "Unknown error"}`);
                        return;
                    }

                    alert("Stock deleted successfully!");
                    window.location.href = "/stocks";
                } catch (error) {
                    alert("An error occurred while deleting the stock.");
                }
            }
        }
    </script>
</body>
</html>
