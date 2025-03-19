// authUtils.js

/**
 * Fetch captcha image and id from the server.
 */
export function fetchCaptcha() {
    return fetch('/user/captcha')
        .then(response => response.json())
        .then(data => {
            return {
                captchaImage: data.captcha_image,
                captchaID: data.captcha_id
            };
        })
        .catch(error => {
            console.error('Error fetching captcha:', error);
            throw error;
        });
}

/**
 * Handle login success by storing the token and redirecting.
 * @param {string} token - The authentication token received from the server.
 * @param {string} redirectUrl - The URL to redirect to after login.
 */
export function handleLoginSuccess(token, redirectUrl) {
    if (token && token.length > 0) {
        localStorage.setItem('Authorization', token);
        vant.Toast.success("Login success!");
        setTimeout(() => {
            window.location.href = redirectUrl;
        }, 2000);
    } else {
        vant.Toast.fail('Login failed. Please try again.');
    }
}

/**
 * Submit login form using Fetch API.
 * @param {Event} event - The form submission event.
 */
export function submitForm(event) {
    event.preventDefault(); // Prevent default form submission

    const formData = new FormData(document.getElementById('loginForm'));

    fetch('/user/login', {
        method: 'POST',
        body: formData,
    })
        .then(response => response.json())
        .then(data => {
            console.log('Login response:', data);
            handleLoginSuccess(data.token, '/user/home');
        })
        .catch(error => {
            console.error('Error:', error);
            vant.Toast.fail('An error occurred. Please try again.');
        });
}

/**
 * Make a fetch request with authentication token.
 * @param {string} url - The URL to fetch from.
 * @param {Object} [options={}] - Optional fetch options.
 */
export function fetchWithAuth(url, options = {}) {
    const token = localStorage.getItem('Authorization');
    console.log('***********************');
    console.log(url);
    if (!token) {
        console.error('No token found in localStorage');
        return Promise.reject('No token found');
    }

    options.headers = {
        ...options.headers,
        'Authorization': `${token}`
    };
    console.log(token);
    console.log(url);
    return fetch(url, options)
        .then(response => response.json())
        .catch(error => {
            console.error('Error:', error);
            throw error;
        });
}
