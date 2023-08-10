import { useRequestRaw } from "../client"
import { create } from 'zustand'

export type Selection = { value: string, label: string }

export type ListProfile = {
    id: string
    emailOrUsername: string
    password: string
    limitUpload: number
    markupName: string
    spinName: string
    colName: string
    productCount: number
    isActive: boolean
    isChecked: boolean
    secret: string
}

export type ListProfileState = {
    totalData: number
    pendingInit: boolean
    clipboard: ListProfile | null
    list: Array<ListProfile>
    spins: Array<Selection>
    markups: Array<Selection>
    collections: Array<Selection>
    error: null | string
}

export type ListProfileActions = {
    initEffect: (limit: number, offset: number, search: string) => void
    updateSingleProfile: (id: string, profile: Partial<ListProfile>) => void
    updateAllProfileWith: (profile: Partial<ListProfile>) => void
    replaceAllProfile: (profiles: Array<ListProfile>) => void
    setClipboard: (profile: ListProfile | null) => void
}

export const useListProfileStore = create<ListProfileState & ListProfileActions>(
    (set, get) => {
        const { sender: getAccountList } = useRequestRaw("GetTokopediaAkunList")
        const { sender: getMarkupList } = useRequestRaw("GetLegacyApiListMarkup")
        const { sender: getSpinList } = useRequestRaw("GetLegacyApiSettingSpin")
        const { sender: getcollectionList } = useRequestRaw("GetLegacyV1ProductNamespaceAll")

        return {
            list: [],
            clipboard: null,
            pendingInit: false,
            collections: [],
            markups: [],
            error: null,
            totalData: 0,
            spins: [],
            initEffect: (limit, offset, search) => {
                set(state => ({ ...state, pendingInit: true, error: null, list: [] }))

                getSpinList({ method: "get", path: "legacy/api/settingSpin" }, {
                    onSuccess(data) {
                        const spins: Selection[] = data.titlePool.map(s => ({ label: s.name, value: s.name }))
                        set(state => ({ ...state, spins }))
                    },
                })

                getMarkupList({ method: "get", path: "legacy/api/listMarkup" }, {
                    onSuccess(data) {
                        const markups: Selection[] = data.data.map(m => ({ label: m, value: m }))
                        set(state => ({ ...state, markups }))
                    },
                })

                getcollectionList({ method: "get", path: "legacy/v1/product/namespace_all" }, {
                    onSuccess(data) {
                        const collections: Selection[] = data.map(d => ({ label: d.name, value: d.name }))
                        set(state => ({ ...state, collections }))
                    },
                })

                getAccountList({ method: "get", path: "tokopedia/akun/list", params: { limit, offset, search } }, {
                    onSuccess: (data) => {
                        const profile = data?.data

                        if (profile) {
                            const profiles = profile.map<ListProfile>(p => ({
                                colName: p.collection,
                                emailOrUsername: p.username,
                                id: p.username,
                                isActive: p.active_upload,
                                limitUpload: p.limit_upload,
                                markupName: p.markup,
                                password: p.password,
                                productCount: p.count_upload,
                                spinName: p.spin,
                                isChecked: false,
                                secret: ""
                            }))
                            set(state => ({ ...state, list: profiles, error: null }))
                        }
                        set(state => ({ ...state, pendingInit: false, error: null, totalData: data.pagination.count }))
                    },
                    onError(e) { set(state => ({ ...state, pendingInit: false, error: e.msg })) }
                })

            },
            updateSingleProfile: (id, profile) => {
                const list = get().list
                const newList = list.map(pro => {
                    if (pro.id == id) {
                        return { ...pro, ...profile }
                    }

                    return pro
                })

                set(state => ({ ...state, list: newList }))
            },
            setClipboard: (profile) => set(state => ({ ...state, clipboard: profile })),
            updateAllProfileWith: (profile) => set(state => ({ ...state, list: state.list.map(current => ({ ...current, ...profile })) })),
            replaceAllProfile: (profiles) => set(state => ({ ...state, list: profiles }))
        }
    }
)