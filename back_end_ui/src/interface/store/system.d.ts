export class SystemInfo {
  public theme: string
  //主题颜色
  public colorPrimary: string
  //背景颜色
  public bgBaseColor: string
  public darkBgBaseColor: string
  //当前语言
  public lang: string
  public layout: string
  public sidePosition: string
  public showSide: boolean
  public selectedKeys : string[]

  constructor() {
    this.theme = 'light'
    this.lang = 'zh-CN'
    this.layout = '1'
    this.sidePosition = 'left'
    this.showSide = true
    this.darkBgBaseColor = '#141414'
    this.bgBaseColor = '#ffffff'
    this.colorPrimary = '#1677ff'
  }
}
