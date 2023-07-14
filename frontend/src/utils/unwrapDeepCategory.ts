import { Category, CategoryAllListLiteRes } from "../client/sdk_types";

type UnwrapDeepCategoryRes = Array<{
    value: number
    label: string
}>

export const unwrapDeepCategory = (data: CategoryAllListLiteRes): UnwrapDeepCategoryRes => {
    const res: UnwrapDeepCategoryRes = [];
    const list = data.data.categoryAllListLite;

    function recurse(c: Category) {
        if (c.children) {
            c.children.forEach(recurse)
        } else {
            res.push({value: c.id, label: c.name})
        }
    }

    if (list) {
        list.categories.forEach(c => {
            if (c.children) {
                c.children.forEach(recurse)
            } else {
                res.push({value: c.id, label: c.name})
            }
        })
    }

    return res;
}