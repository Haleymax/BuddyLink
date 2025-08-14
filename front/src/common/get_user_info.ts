import { useAuthStore } from '../stores/auth.store'
import { getUserInfo } from '../api/auth'

type Notify = {
    error?: (msg: string) => void
}

export function useFetchUserInfo() {
    const authStore = useAuthStore()

    const userData = (data: any) => ({
        id: data.id,
        uuid: data.uuid,
        email: data.email,
        username: data.username,
        avatar: data.avatar || null,
        role: data.role
    })

    /**
     * 通用获取并同步用户信息
     * @param notify 可传入 naive-ui 的 message 实例: { error: message.error }
     * @returns { ok: boolean, data?: any, error?: unknown }
     */

    const fetchUserInfoOnce = async (notify?: Notify) => {
        const token = authStore.token
        console.log('fetchUserInfoOnce called, token:', token ? 'exists' : 'missing')
        
        if (!token) {
            console.log('No token found')
            notify?.error?.('No token found, please login')
            return { ok: false }
        }
        
        try {
            console.log('Calling getUserInfo API...')
            const response = await getUserInfo(token)
            console.log('getUserInfo response:', response)
            
            if (response.code !== 200) {
                console.log('API response status is not 200:', response.code)
                notify?.error?.('Failed to fetch user info, please try login again')
                return { ok: false }
            }
            
            const user = userData(response.data)
            console.log('Processed user data:', user)
            authStore.user = user
            return { ok: true, data: user }
        } catch (error) {
            console.error('Error fetching user info:', error)
            notify?.error?.('Failed to fetch user info, please try login again')
            return { ok: false, error}
        }
    }

    return { fetchUserInfoOnce }
}