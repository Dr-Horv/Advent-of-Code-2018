package day07

import (
	"fmt"
	"github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"sort"
	"strings"
)

const OFFSET = -('A' - 1)

type node struct {
	Id       string
	parents  []node
	children []node
	Work     int
}

func getNodeOrCreate(id string, nodes map[string]node) node {
	n, found := nodes[id]

	if !found {
		work := ([]rune(id))[0] + OFFSET + 60
		return node{id, make([]node, 0), make([]node, 0), int(work)}
	}

	return n
}

func Solve(lines []string, partOne bool) string {

	nodes := make(map[string]node)
	for _, l := range lines {
		parentId, childId := parseLine(l)
		parent := getNodeOrCreate(parentId, nodes)
		child := getNodeOrCreate(childId, nodes)

		parent.children = append(parent.children, child)
		child.parents = append(child.parents, parent)

		nodes[parentId] = parent
		nodes[childId] = child
	}

	var answer = ""
	if partOne {
		for {
			readyNodes := make([]node, 0)
			for _, n := range nodes {
				if len(n.parents) == 0 {
					readyNodes = append(readyNodes, n)
				}
			}

			if len(readyNodes) == 0 {
				return answer
			}

			sort.Slice(readyNodes, func(i, j int) bool {
				return readyNodes[i].Id < readyNodes[j].Id
			})

			nextNode := readyNodes[0]

			answer += nextNode.Id
			complete(nextNode, nodes)
			delete(nodes, nextNode.Id)
		}
	}

	nodeAssignments := make(map[int]node)
	workerAssignments := make(map[string]int)
	availableWorkers := []int{1, 2, 3, 4, 5}
	second := 0
	for {
		if len(availableWorkers) > 0 {
			readyNodes := make([]node, 0)
			for _, n := range nodes {
				_, found := workerAssignments[n.Id]
				if !found && len(n.parents) == 0 {
					readyNodes = append(readyNodes, n)
				}
			}

			sort.Slice(readyNodes, func(i, j int) bool {
				return readyNodes[i].Id < readyNodes[j].Id
			})

			if len(readyNodes) > 0 {
				bound := pkg.Min(len(availableWorkers), len(readyNodes))
				for i := 0; i < bound; i++ {
					nodeAssignments[availableWorkers[0]] = readyNodes[i]
					workerAssignments[readyNodes[i].Id] = availableWorkers[0]
					availableWorkers = availableWorkers[1:]
				}
			}
		}

		if len(nodeAssignments) == 0 {
			return fmt.Sprint(second)
		}

		for w, n := range nodeAssignments {
			n.Work = n.Work - 1
			nodeAssignments[w] = n

			if n.Work == 0 {
				answer += n.Id
				complete(n, nodes)
				delete(nodes, n.Id)
				availableWorkers = append(availableWorkers, w)
				delete(nodeAssignments, w)
				delete(workerAssignments, n.Id)
			}
		}

		second++
	}

}

func removeNode(n node, ns []node) []node {
	for i, curr := range ns {
		if curr.Id == n.Id {
			return append(ns[:i], ns[i+1:]...)
		}
	}

	return ns
}

func complete(n node, nodes map[string]node) {
	for id, curr := range nodes {
		curr.parents = removeNode(n, curr.parents)
		nodes[id] = curr
	}
}

func parseLine(s string) (parent string, child string) {
	parts := strings.Split(s, " ")
	parent = strings.TrimSpace(parts[1])
	child = strings.TrimSpace(parts[7])

	return parent, child
}
