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
    // 檢查認證
    if (!checkAuth()) return;

    const urlParams = new URLSearchParams(window.location.search);
    const productId = urlParams.get('id');
    const productName = decodeURIComponent(urlParams.get('name') || '');

    // 填充產品名稱
    const productNameInput = document.getElementById('productName');
    productNameInput.value = productName;

    const submitBtn = document.getElementById('submitBtn');
    submitBtn.addEventListener('click', async () => {
        const price = document.getElementById('price').value;
        const memo = document.getElementById('memo').value;

        if (!price) {
            alert('請輸入售價');
            return;
        }

        const saleData = {
            productId: parseInt(productId, 10),
            price: parseFloat(price),
            memo: memo || ''
        };

        try {
            const response = await fetch(`${BASE_URL}/sale`, {
                method: 'POST',
                headers: getAuthHeaders(),
                body: JSON.stringify(saleData)
            });

            if (!response.ok) {
                if (response.status === 401) {
                    localStorage.removeItem('token');
                    window.location.href = 'login.html';
                    return;
                }
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            alert('上架成功！');
            window.location.href = 'index.html'; // 成功後返回首頁
        } catch (error) {
            console.error('Error submitting sale:', error);
            alert('上架失敗，請稍後再試');
        }
    });
});