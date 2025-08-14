import { defineStore } from 'pinia';

export interface User {
    id?: number;
    uuid?: string;
    email?: string;
    username?: string;
    avatar?: string | null;
    role?: string;
}

interface AuthState {
    user: User | null;
    token: string | null;
}

export const useAuthStore = defineStore('auth', {
    state: (): AuthState => ({
        user: null,
        token: null,
    }),
    getters: {
        isAuthenticated: (state) => !!state.token,
        getToken: (state):string | null => state.token,
        getUser: (state): User | null => state.user,
    },
    actions: {
        SetToken(token: string | null): void {
            this.token = token;
            if (token) {
                localStorage.setItem('auth_token', token);
            } else {
                localStorage.removeItem('auth_token');
            }
        },

        SetUser(user: User | null): void {
            this.user = user;
            if (user) {
                localStorage.setItem('auth_user', JSON.stringify(user));
            } else {
                localStorage.removeItem('auth_user');
            }
        },

        loadFormStorage(): void {
            const token = localStorage.getItem('auth_token');
            const user = localStorage.getItem('auth_user');

            if (token) {
                this.SetToken(token);
            }
            if (user) {
                try {
                    this.SetUser(JSON.parse(user));
                } catch (e) {
                    console.error('Failed to parse user from storage:', e);
                    this.SetUser(null);
                }
            }
        },
        clearAuth(): void {
            this.token = null
            this.user = null
            localStorage.removeItem('auth_token');
            localStorage.removeItem('auth_user');
        }
    }
}); 