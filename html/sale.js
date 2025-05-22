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
    if (!checkAuth()) return;

    const urlParams = new URLSearchParams(window.location.search);
    const productId = urlParams.get('id');
    let productName = '未提供名稱';

    // 除錯：打印原始 name 參數
    const nameParam = urlParams.get('name');
    console.log('Raw name param:', nameParam);

    // 嘗試解碼 name 參數
    try {
        productName = nameParam ? decodeURIComponent(nameParam) : '未提供名稱';
        console.log('Decoded product name:', productName); // 除錯：確認解碼結果
    } catch (error) {
        console.error('Error decoding name parameter:', error, 'Raw param:', nameParam);
        alert('產品名稱格式錯誤，請從產品列表重新進入');
        window.location.href = 'index.html';
        return;
    }

    // 填充產品名稱
    const productNameInput = document.getElementById('productName');
    if (!productNameInput) {
        console.error('productName input not found');
        alert('頁面載入錯誤，請稍後再試');
        return;
    }
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
            memo: memo || ' '
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
            window.location.href = 'index.html';
        } catch (error) {
            console.error('Error submitting sale:', error);
            alert('上架失敗，請稍後再試');
        }
    });
});