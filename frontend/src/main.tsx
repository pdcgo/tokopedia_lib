import ReactDOM from "react-dom/client"
import App from "./App.tsx"
import { ConfigProvider, ThemeConfig } from "antd"
import "./index.scss"
import "antd/dist/reset.css"

const theme: ThemeConfig = {
    token: {
        fontFamily:
            "Roboto, -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Helvetica Neue', Arial, 'Noto Sans', sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol', 'Noto Color Emoji' ",
        fontSize: 13.4,
        colorPrimary: "#FF6000",
        colorInfo: "#FF6000",
        colorError: "#e7b072",
        colorErrorHover: "#e0732b",
        colorErrorBorder: "#FFE6C7",
        colorErrorBorderHover: "#d88b34",
        colorText: "#454545",
    },
    components: {
        Tooltip: {
            colorBgDefault: "#FF6000",
            fontSize: 13,
            colorBorder: "#FF6000",
            sizePopupArrow: 12,
        },
        Pagination: {
            colorBgContainer: "#fefffece",
        },
        Input: {
            fontSize: 13.4,
            fontSizeLG: 13.5,
        },
        Button: {
            boxShadow: "none",
            boxShadowSecondary: "none",
            boxShadowTertiary: "none",
        },
        Card: {
            colorBorderSecondary: "#e1dfdd",
        },
    },
}

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
    <ConfigProvider prefixCls="nox" theme={theme}>
        <App />
    </ConfigProvider>
)
