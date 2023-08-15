import type {IStatConfigResp, IStatInfoResp} from "$lib/models";

export const getStatiInfo = async (): Promise<IStatInfoResp> => {
    return new Promise<IStatInfoResp>(async (resolve, reject) => {
        await fetch("http://localhost:3000/debug/stati/info").then(async resp => {
            let body = await resp.json()
            resolve(body)
        }).catch(e => reject(e))
    })
}

export const getStatiConfig = async (): Promise<IStatConfigResp> => {
    return new Promise<IStatConfigResp>(async (resolve, reject) => {
        await fetch("http://localhost:3000/debug/stati/config").then(async resp => {
            let body = await resp.json()
            resolve(body)
        }).catch(e => reject(e))
    })
}