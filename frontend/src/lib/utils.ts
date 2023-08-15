/**
 * Okunabilir şekilde byte'nın gösterilmesi
 *
 * @param bytes Number olarak byte miktarı.
 * @param si SI metrik olarak kulanılırsa 1000'e bölünmesi gerekir, Eğer SI metrik değilse (IEC) 1024'e bölünür
 * @param dp gösterilmesi istenilen sayının küsürat uzunluğu.
 *
 * @return okunabilir byte bilgisi.
 */
export function humanFileSize(bytes: number, si = true, dp = 1): string {
    const thresh = si ? 1000 : 1024;
    if (Math.abs(bytes) < thresh) {
        return bytes + " B";
    }
    const units = si
        ? ["kB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"]
        : ["KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB"];
    let u = -1;
    const r = 10 ** dp;
    do {
        bytes /= thresh;
        ++u;
    } while (
        Math.round(Math.abs(bytes) * r) / r >= thresh &&
        u < units.length - 1
        );
    return bytes.toFixed(dp) + " " + units[u];
}