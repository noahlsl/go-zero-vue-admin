package menu

import (
	"zero/internal/types"
)

// buildMenuTree 将菜单列表组装为树形结构
func buildMenuTree(menus []types.SysMenu) []types.SysMenu {
	treeMap := make(map[uint][]types.SysMenu)
	for _, m := range menus {
		treeMap[m.ParentId] = append(treeMap[m.ParentId], m)
	}
	roots := treeMap[0]
	for i := 0; i < len(roots); i++ {
		attachChildren(&roots[i], treeMap)
	}
	return roots
}

func attachChildren(menu *types.SysMenu, treeMap map[uint][]types.SysMenu) {
	menu.Children = treeMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		attachChildren(&menu.Children[i], treeMap)
	}
}
