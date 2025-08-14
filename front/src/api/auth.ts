import request from '../utils/request';
import type { LoginFormData, RegisterFormData } from '../model/auth';

export const login = (data: LoginFormData) => {
    return request({
        url: '/api/v1/user/login',
        method: 'post',
        data: data
    })
}

export const register = (data: RegisterFormData) => {
    const formData = new FormData();
    formData.append('email', data.email);
    formData.append('username', data.username);
    formData.append('password', data.password);
    formData.append('verification', data.verification_code);

    return request({
        url: '/api/v1/user/register',
        method: 'post',
        data: formData,
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    })
}


export const sendVerificationCode = (email: string) => {
    const formData = new FormData();
    formData.append('email', email);

    return request({
        url: '/api/v1/user/send_captcha',
        method: 'post',
        data: formData,
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    })
}

export const getUserInfo = (token: string) => {
    return request({
        url: '/api/v1/user/get_user',
        method: 'get',
        headers: {
            "Authorization": token,
        }
    })
}
