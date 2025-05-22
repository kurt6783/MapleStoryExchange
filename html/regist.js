const BASE_URL = 'https://www.maplestoryexchange.com';
// const BASE_URL = 'http://localhost:8080';

document.addEventListener('DOMContentLoaded', () => {
    const registerForm = document.getElementById('registerForm');
    const message = document.getElementById('message');

    registerForm.addEventListener('submit', async (e) => {
        e.preventDefault();

        const account = document.getElementById('account').value.trim();
        const password = document.getElementById('password').value;
        const password_confirm = document.getElementById('password_confirm').value;
        const code = document.getElementById('code').value.trim();

        // 驗證密碼一致性
        if (password !== password_confirm) {
            message.textContent = '密碼與確認密碼不一致';
            message.classList.add('error');
            return;
        }

        const payload = {
            account,
            password,
            password_confirm,
            code
        };

        try {
            const response = await fetch(`${BASE_URL}/regist`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(payload)
            });

            const data = await response.json();
            if (response.ok) {
                message.textContent = data.message || '註冊成功！即將跳轉到登入頁面...';
                message.classList.remove('error');
                registerForm.reset(); // 清空表單
                // 延遲 2 秒後跳轉到登入頁面
                setTimeout(() => {
                    window.location.href = 'login.html';
                }, 2000);
            } else {
                message.textContent = data.error || '註冊失敗';
                message.classList.add('error');
            }
        } catch (error) {
            console.error('Register error:', error);
            message.textContent = '註冊失敗，請檢查網路';
            message.classList.add('error');
        }
    });
});