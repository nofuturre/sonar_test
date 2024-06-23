import React from 'react';
import useFetch from '../hooks/useFetch';

const pay = (productId) => {
    fetch("http://localhost:8080/products/pay", {
        method: 'POST',
        mode: 'no-cors',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ id: productId }),
    })
        .catch(error => {
            console.error('Error sending product ID to server:', error);
        });
    window.location.reload();
};

const Products = () => {
    const {data: products, loading, error} = useFetch("http://localhost:8080/products");

    if (loading) return <div>Loading...</div>;
    if (error) return <div>Error: {error}</div>;

    return (
        <div>
            <h2>Products</h2>
            <table>
                <thead>
                <tr>
                    <th>ID</th>
                    <th>Name</th>
                    <th>Price</th>
                    <th>Paid</th>
                </tr>
                </thead>
                <tbody>
                {products.map(product => (
                    <tr key={product.id}>
                        <td>{product.id}</td>
                        <td>{product.name}</td>
                        <td>{product.price}</td>
                        <td>{product.paid ? "✓" : ""}</td>
                        <td>
                            <button onClick={() => pay(product.id)}>Pay</button>
                        </td>
                    </tr>

                ))}
                </tbody>
            </table>
        </div>
    );
};

export default Products;
