import request from '../utils/request';
import type { LoginData, RegisterFormData } from '../model/auth';

export const login = (data: LoginData) => {
    return request({
        url: '/api/v1/user/login',
        method: 'post',
        data: data
    })
}

// 用于注册用户的接口 使用的是表单数据
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


// 用于发送验证码的接口 使用的是表单数据
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