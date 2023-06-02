/* eslint-disable @typescript-eslint/no-non-null-assertion */
import {
    ApiOutlined,
    CloudUploadOutlined,
    ShopOutlined,
} from "@ant-design/icons"
import { Button, Card, Tabs, Tooltip } from "antd"
import { useEffect, useState } from "react"
import {
    AppContainer,
    AppContainer2,
    FlexColumn,
    FloatingMenu,
} from "./styled_components"
import AddAccount from "./views/AddAccount"
import CategoryMapping from "./views/CategoryMapping"
import Upload from "./views/Upload"

const menus = [
    {
        key: "add_account",
        name: "Add Account",
        icon: ShopOutlined,
        child: AddAccount,
    },
    {
        key: "upload",
        name: "Upload",
        icon: CloudUploadOutlined,
        child: Upload,
    },
    {
        key: "category_map",
        name: "Category Mapping",
        icon: ApiOutlined,
        child: CategoryMapping,
    },
]

function App() {
    const [activeMenu, setActiveMenu] = useState("add_account")
    const [showFloatMenu, setShowFloatMenu] = useState(false)

    useEffect(() => {
        const observer = new IntersectionObserver(
            (entry) => {
                if (!entry[0].isIntersecting) setShowFloatMenu(true)
                else setShowFloatMenu(false)
            },
            { threshold: [0] }
        )

        observer.observe(document.getElementById("add_account")!)

        return () => {
            observer.unobserve(document.getElementById("add_account")!)
        }
    }, [])

    return (
        <AppContainer>
            <div id="top"></div>
            <AppContainer2>
                <FloatingMenu show={showFloatMenu}>
                    <Card size="small" title="Page" type="inner">
                        <FlexColumn>
                            {menus.map((menu) => (
                                <Tooltip
                                    key={menu.key}
                                    title={menu.name}
                                    placement="right"
                                >
                                    <Button
                                        shape="circle"
                                        key={menu.key}
                                        type={
                                            activeMenu == menu.key
                                                ? "primary"
                                                : "default"
                                        }
                                        onClick={() => setActiveMenu(menu.key)}
                                        icon={<menu.icon rev={menu.key} />}
                                    ></Button>
                                </Tooltip>
                            ))}
                        </FlexColumn>
                    </Card>
                </FloatingMenu>
                <Tabs
                    style={{ height: "100%" }}
                    activeKey={activeMenu}
                    onChange={setActiveMenu}
                    items={menus.map((menu) => ({
                        key: menu.key,
                        label: (
                            <span
                                id={menu.key}
                                style={{ fontWeight: 400, userSelect: "none" }}
                            >
                                <menu.icon rev={menu.key} />
                                {menu.name}
                            </span>
                        ),
                        children: (
                            <menu.child
                                key={activeMenu}
                                activePage={activeMenu}
                            />
                        ),
                    }))}
                />
            </AppContainer2>
        </AppContainer>
    )
}

export default App
