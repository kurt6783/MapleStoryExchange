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

    const myList = document.getElementById('myList');
    let mys = [];

    // 從 API 獲取資料
    async function fetchMys() {
        try {
            const response = await fetch(`${BASE_URL}/my`, {
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
            mys = data.data ?? []; // 儲存 API 回傳的資料

            displayMys(mys); // 初始顯示所有產品
        } catch (error) {
            console.error('Error fetching mys:', error);
            myList.innerHTML = '<tr><td colspan="4">無法載入資料，請檢查 API 或登入狀態</td></tr>';
        }
    }

    // 顯示產品列表
    function displayMys(mys) {
        myList.innerHTML = '';
        if (mys.length === 0) {
            myList.innerHTML = '<tr><td colspan="4">無匹配結果</td></tr>';
            return;
        }
        mys.forEach(my => {
            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${my.name}</td>
                <td>${Number(item.price).toLocaleString()}</td>
                <td>${my.memo}</td>
                <td><button onclick="removeProduct('${my.id}')">下架產品</button></td>
            `;
            myList.appendChild(row);
        });
    }

    // 下架產品
    async function removeProduct(id) {
        if (!confirm('確定要下架此產品嗎？')) return;

        try {
            const response = await fetch(`${BASE_URL}/remove`, {
                method: 'POST',
                headers: getAuthHeaders(),
                body: JSON.stringify({ product_id: parseInt(id, 10) })
            });

            if (!response.ok) {
                if (response.status === 401) {
                    localStorage.removeItem('token');
                    window.location.href = 'login.html';
                    return;
                }
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            alert('下架成功！');
            fetchMys(); // 重新載入產品列表
        } catch (error) {
            console.error('Error removing product:', error);
            alert('下架失敗，請稍後再試');
        }
    }

    // 初始化載入資料
    fetchMys();

    // 將 removeProduct 函數暴露到全局作用域，以便 onclick 使用
    window.removeProduct = removeProduct;
});