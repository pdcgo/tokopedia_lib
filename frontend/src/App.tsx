/* eslint-disable @typescript-eslint/no-non-null-assertion */
import {
    ApiOutlined,
    CloudUploadOutlined,
    UsergroupAddOutlined,
    CodeSandboxOutlined,
} from "@ant-design/icons"
import { Button, Card, Tabs, Tooltip } from "antd"
import React, { Suspense, useEffect, useState } from "react"
import {
    AppContainer,
    AppContainer2,
    FlexColumn,
    FloatingMenu,
} from "./styled_components"

const EtalaseMapping = React.lazy(() => import("./views/EtalaseMapping"))
const AddAccount = React.lazy(() => import("./views/AddAccount"))
const CategoryMapping = React.lazy(() => import("./views/CategoryMapping"))
const Upload = React.lazy(() => import("./views/Upload"))

const menus = [
    {
        key: "accounts",
        name: "Accounts",
        icon: UsergroupAddOutlined,
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
    {
        key: "etalase_map",
        name: "Etalase Mapping",
        icon: CodeSandboxOutlined,
        child: EtalaseMapping,
    },
]

function App() {
    const [activeMenu, setActiveMenu] = useState("accounts")
    const [showFloatMenu, setShowFloatMenu] = useState(false)

    useEffect(() => {
        const element = document.getElementById("accounts")
        const observer = new IntersectionObserver(
            (entry) => {
                if (!entry[0].isIntersecting) setShowFloatMenu(true)
                else setShowFloatMenu(false)
            },
            { threshold: [0] }
        )

        if (element) {
            observer.observe(element)
        }

        return () => {
            if (element) {
                observer.unobserve(element)
            }
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
                            <Suspense
                                key={activeMenu}
                                fallback={<Card loading />}
                            >
                                <menu.child
                                    key={activeMenu}
                                    activePage={activeMenu}
                                />
                            </Suspense>
                        ),
                    }))}
                />
            </AppContainer2>
        </AppContainer>
    )
}

export default App
