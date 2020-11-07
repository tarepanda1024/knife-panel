
export const  formatSize = (size) => {
    let units = ['B', 'KB', 'M', 'G', 'T', 'P', 'E'];
    let unitIndex = 0;
    while (size >= 1024 && unitIndex < units.length - 1) {
        size = size / 1024;
        unitIndex++;
    }
    return size.toFixed(1) + units[unitIndex]
}