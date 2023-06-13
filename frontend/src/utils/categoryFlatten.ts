import { Category } from "../client/sdk_types"

export function categoryFlatten(c?: Category[]) {
    const callback = (penampung: number[][], parentIndexs: number[], c: Category) => {
        if (!c.children) {
            penampung.push([...parentIndexs, c.id])
        } else {
            c.children.forEach((chil) =>
                callback(penampung, [...parentIndexs, c.id], chil)
            )
        }
    }
    const penampung: number[][] = []
    c?.forEach((cat) => callback(penampung, [], cat))
    return penampung
}