/*
Package flood_fill
https://leetcode-cn.com/problems/flood-fill/

733. 图像渲染

有一幅以 m x n 的二维整数数组表示的图画 image ，其中 image[i][j] 表示该图画的像素值大小。

你也被给予三个整数 sr ,  sc 和 newColor 。你应该从像素 image[sr][sc] 开始对图像进行 上色填充 。

为了完成 上色工作 ，从初始像素开始，记录初始坐标的 上下左右四个方向上 像素值与初始坐标相同的相连像素点，
接着再记录这四个方向上符合条件的像素点与他们对应 四个方向上 像素值与初始坐标相同的相连像素点，……，重复该过程。
将所有有记录的像素点的颜色值改为 newColor 。

最后返回 经过上色渲染后的图像 。

*/
package flood_fill

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	type point struct {
		x, y int
	}
	points := []point{point{sr, sc}}
	dirs := []int{1, 0, -1, 0, 1}
	m, n := len(image), len(image[0])
	oldColor := image[sr][sc]
	visited := make(map[point]bool)
	visited[point{sr, sc}] = true
	for len(points) > 0 {
		size := len(points)
		for i := 0; i < size; i++ {
			p := points[i]
			image[p.x][p.y] = newColor
			for i := 0; i < len(dirs)-1; i++ {
				x, y := p.x+dirs[i], p.y+dirs[i+1]
				if x >= 0 && x < m && y >= 0 && y < n && image[x][y] == oldColor && !visited[point{x, y}] {
					points = append(points, point{x, y})
					visited[point{x, y}] = true
				}
			}
		}
		points = points[size:]
	}
	return image
}
