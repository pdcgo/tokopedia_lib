type SetCacheData = {
    email: string
    password: string
    secret: string
}

export function cacheStorage() {
    const TOPED_EMAIL_FOR_GET_CAT = "email_ctx_toped"
    const TOPED_PASSWORD_FOR_GET_CAT = "password_ctx_toped"
    const TOPED_SECRET_FOR_GET_CAT = "secret_ctx_toped"

    function getCache() {
        const email = localStorage.getItem(TOPED_EMAIL_FOR_GET_CAT)
        const password = localStorage.getItem(TOPED_PASSWORD_FOR_GET_CAT)
        const secret = localStorage.getItem(TOPED_SECRET_FOR_GET_CAT)

        if (email && password && secret) return { email, password, secret }
        return null
    }

    function setCache(data: SetCacheData) {
        localStorage.setItem(TOPED_EMAIL_FOR_GET_CAT, data.email)
        localStorage.setItem(TOPED_PASSWORD_FOR_GET_CAT, data.password)
        localStorage.setItem(TOPED_SECRET_FOR_GET_CAT, data.secret)
    }

    return { getCache, setCache }
}