import { create } from "zustand"
import { useRequestRaw } from "../client"
import { CategoryAllListLiteRes } from "../client/sdk_types"
import { categoryFlatten } from "../utils/categoryFlatten"

/**@description Digunakan pada komponen `view/CategoryMapping.tsx`*/
export type ListMapper = {
    /** Digunakan untuk komponen `Breadcrumb` pada komponen `MapCard`
     * @example ["Rumah Tangga", "Dekorasi", "Cover Kipas Angin"]*/
    shopeeCategoryName: string[]
    /** Menunjukan jumlah produk total yang ada pada kategori terkait 
     * @description Digunakan di komponen `MapCard` sebelah kanan atas */
    productCount: number
    /** Menunjukan kategori asli produk di shopee */
    shopeeCategoryId: number
    /** Array dari target mapping kategori & berisi id kategori dari Tokopedia. 
     * @description Harus berurutan dari id kategori teratas ke yang palig bawah.
     * @example [1233, 4566, 674]*/
    tokopediaCategoryIds: (number | string)[]
}

export type ListMapperState = {
    list: Array<ListMapper>
    pendingInitEffect: boolean
}

export type ListMapperActions = {
    /** Digunakan untuk mendapatkan data awal mapping.
     * @tutorial Gunakan dalam fungsi `React.useEffect` */
    initEffect: (namespace: string, topedCategories: CategoryAllListLiteRes | null) => void
    /** Digunakan untuk update salah satu data list.
     * @description Gunakan ketika list update, contoh -> `Saat update tokopediaCategoryIds` */
    updateSingleList: (shopeeCategoryId: number, change?: Partial<ListMapper>) => void
}

export const useListStore = create<ListMapperState & ListMapperActions>(
    (set) => {
        const {
            /** Digunakan untuk mendapatkan data awal mapping dari koleksi berdasarkan nama koleksi */
            sender: getInitialMapData
        } = useRequestRaw("GetV1ProductCategory")

        const {
            /** Digunakan setelah mendapatkan data awal map kemudian mencocokan tokopedia id jika ada */
            sender: getMapAfterInitialData
        } = useRequestRaw("GetTokopediaMapperMap")

        return {
            list: [],
            pendingInitEffect: false,
            initEffect: (namespace, topedCategories) => {
                set(s => ({ ...s, pendingInitEffect: true, list: [] }))
                // dapatkan data dulu dari koleksi terkait
                getInitialMapData({
                    method: "get",
                    path: "v1/product/category",
                    params: {
                        is_public: false,
                        marketplace: "shopee",
                        namespace,
                        kota: "",
                        pmax: 0,
                        pmin: 0
                    }
                }, {
                    onError: () => { set(s => ({ ...s, pendingInitEffect: false })) },
                    onSuccess: (data) => {
                        if (data) {
                            if (!data.length) set(s => ({ ...s, pendingInitEffect: false, list: [] }))

                            const newList = data.map<ListMapper>(map => ({ productCount: map.count, shopeeCategoryId: map._id, shopeeCategoryName: map.name, tokopediaCategoryIds: [] }))
                            set(s => ({ ...s, pendingInitEffect: false, list: newList }))

                            // terus dapatkan data map dari hasil data di atas
                            getMapAfterInitialData({
                                method: "get",
                                path: "tokopedia/mapper/map",
                                params: {
                                    collection: namespace,
                                }
                            }, {
                                onSuccess(data) {
                                    if (topedCategories) {
                                        const flattenCats = categoryFlatten(topedCategories.data.categoryAllListLite?.categories)

                                        data.data.forEach((rdata) => {
                                            flattenCats.forEach((fc) => {
                                                if (fc.indexOf(rdata.tokopedia_id) > -1) {
                                                    set(state => ({
                                                        ...state,
                                                        list: state.list.map((ls) => {
                                                            if (ls.shopeeCategoryId === rdata.shopee_id) ls.tokopediaCategoryIds = fc
                                                            return ls
                                                        })
                                                    }))
                                                }
                                            })
                                        })
                                    }
                                },
                            })
                        }
                    }
                })
            },
            updateSingleList: (shopeeCategoryId, change) => {
                set(state => ({
                    ...state,
                    list: state.list.map(list => {
                        if (list.shopeeCategoryId == shopeeCategoryId) {
                            return { ...list, ...change }
                        }

                        return list
                    })
                }))
            },
        }
    }
)