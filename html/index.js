const BASE_URL = 'http://3.86.236.95:8080';
// const BASE_URL = 'http://localhost:8080';

// 檢查 token 是否存在，若無則跳轉到 login.html
function checkAuth() {
    const token = localStorage.getItem('token');
    if (!token) {
        window.location.href = 'login.html';
        return false;
    }
    return true;
}

// 獲取帶有 Authorization Header 的 fetch 配置
function getAuthHeaders() {
    const token = localStorage.getItem('token');
    return {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
    };
}

document.addEventListener('DOMContentLoaded', () => {
    // 檢查認證，若無 token 則導向 login.html
    if (!checkAuth()) return;

    const searchInput = document.getElementById('searchInput');
    const productList = document.getElementById('productList');
    let products = [];
    let fuse = null;

    // 從 API 獲取資料
    async function fetchProducts() {
        try {
            const response = await fetch(`${BASE_URL}/product`, {
                method: 'GET',
                headers: getAuthHeaders()
            });
            if (!response.ok) {
                if (response.status === 401) {
                    localStorage.removeItem('token');
                    window.location.href = 'login.html';
                    return;
                }
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const data = await response.json();
            products = data.data; // 儲存 API 回傳的資料
            console.log('Fetched products:', products); // 除錯：確認資料
            // 初始化 Fuse.js
            fuse = new Fuse(products, {
                keys: ['Name', 'Category'],
                threshold: 0.2, // 降低門檻以提高中文匹配率
                includeScore: true,
                useExtendedSearch: true,
                tokenize: true,
                matchAllTokens: false
            });
            displayProducts(products); // 初始顯示所有產品
        } catch (error) {
            console.error('Error fetching products:', error);
            productList.innerHTML = '<tr><td colspan="2">無法載入資料，請檢查 API 或登入狀態</td></tr>';
        }
    }

    // 顯示產品列表
    function displayProducts(products) {
        productList.innerHTML = '';
        if (products.length === 0) {
            productList.innerHTML = '<tr><td colspan="3">無匹配結果</td></tr>';
            return;
        }
        products.forEach(product => {
            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${product.name}</td>
                <td><a href="item.html?productId=${product.id}">${product.count}</a></td>
                <td><button onclick="window.location.href='sale.html?id=${product.id}&name=${encodeURIComponent(product.name)}'">我要上架</button></td>
            `;
            productList.appendChild(row);
        });
    }

    // 搜尋功能
    searchInput.addEventListener('input', (e) => {
        const query = e.target.value.trim();
        console.log('Search query:', query); // 除錯：確認輸入值
        if (query && fuse) {
            const results = fuse.search(query);
            console.log('Search results:', results); // 除錯：確認搜尋結果
            displayProducts(results.map(result => result.item));
        } else {
            displayProducts(products); // 無搜尋條件或 fuse 未初始化時顯示所有產品
        }
    });

    // 初始化載入資料
    fetchProducts();
});