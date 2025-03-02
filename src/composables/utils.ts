export function friendlyTheme(id: string): string {
    return id.replaceAll('-', ' ').replace(/(?:^\w)|(?<= )\w/g, w => w.toUpperCase())
}
