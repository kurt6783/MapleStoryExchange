const BASE_URL = 'https://www.maplestoryexchange.com';
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

    const itemList = document.getElementById('itemList');
    let items = [];
    let fuse = null;

    // 從 URL 獲取 productId
    const urlParams = new URLSearchParams(window.location.search);
    const productId = urlParams.get('productId');

    // 從 API 獲取資料
    async function fetchItems() {
        try {
            // 構建 API URL，附加 productId 查詢參數（如果存在）
            const url = productId ? `${BASE_URL}/item?productId=${encodeURIComponent(productId)}` : `${BASE_URL}/item`;
            const response = await fetch(url, {
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
            items = data.data; // 儲存 API 回傳的資料
            console.log('Fetched items:', items); // 除錯：確認資料
            // 初始化 Fuse.js
            fuse = new Fuse(items, {
                keys: ['name', 'category'],
                threshold: 0.2, // 降低門檻以提高中文匹配率
                includeScore: true,
                useExtendedSearch: true,
                tokenize: true,
                matchAllTokens: false
            });
            displayItems(items); // 初始顯示所有產品
        } catch (error) {
            console.error('Error fetching items:', error);
            itemList.innerHTML = '<tr><td colspan="4">無法載入資料，請檢查 API 或登入狀態</td></tr>';
        }
    }

    // 顯示產品列表
    function displayItems(items) {
        itemList.innerHTML = '';
        if (items.length === 0) {
            itemList.innerHTML = '<tr><td colspan="4">無匹配結果</td></tr>';
            return;
        }
        items.forEach(item => {
            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${item.name}</td>
                <td>${item.price}</td>
                <td>${item.memo}</td>
                <td>${item.owner_code}</td>
            `;
            itemList.appendChild(row);
        });
    }

    // 初始化載入資料
    fetchItems();
});