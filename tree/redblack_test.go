package tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedBlackt(t *testing.T) {
	k1, k2, k3, k4, k5, k6, k7, k8 := 1, 2, 3, 4, 5, 6, 7, 8
	v1, v2, v3, v4, v5, v6, v7, v8 := 10, 20, 30, 40, 50, 60, 70, 80
	tree := NewRbTree()
	var info string

	//  k1(b:0)
	tree.Insert(k1, v1)
	info = "{1,10,2,0}"
	assert.Equal(t, tree.levelInfo(), info, fmt.Sprintf("tree.levelInfo() != %s", info))

	//  k1(b:0)
	//	 \
	//    k2(r:-1)
	tree.Insert(k2, v2)
	info = "{1,10,2,0}\n{2,20,1,-1}"
	assert.Equal(t, tree.levelInfo(), info, fmt.Sprintf("tree.levelInfo() != %s", info))
	assert.Equal(t, tree.Height(), 0, "tree.Height() != 0")

	//    	    k2(b:0)
	//     	    /    \
	//     k1(r:-1)   k3(r:-1)
	tree.Insert(k3, v3)
	info = "{2,20,2,0}\n{1,10,1,-1}{3,30,1,-1}"
	assert.Equal(t, tree.levelInfo(), info, fmt.Sprintf("tree.levelInfo() != %s", info))

	//    	    k2(r:0)
	//     	    /    \
	//     k1(r:-1)  k3(r:-1)
	//   			   \
	//   			  k4(r:-1)

	//    	    k2(b:1)
	//     	    /    \
	//     k1(b:0)  k3(b:0)
	//   			   \
	//   			  k4(r:-1)
	tree.Insert(k4, v4)
	info = "{2,20,2,1}\n{1,10,2,0}{3,30,2,0}\n{4,40,1,-1}"
	assert.Equal(t, tree.levelInfo(), info, fmt.Sprintf("tree.levelInfo() != %s", info))
	assert.Equal(t, tree.Height(), 1, "tree.Height() != 1")

	//    	    k2(b:1)
	//     	    /    \
	//     k1(b:0)  k3(b:0)
	//   			   \
	//   			  k4(r:-1)
	//   			     \
	//   			    k5(r:-1)

	//    	    k2(b:1)
	//     	    /    \
	//     k1(b:0)  k4(b:0)
	//   			  / \
	//   	   k3(r:-1)  k5(r:-1)
	tree.Insert(k5, v5)
	info = "{2,20,2,1}\n{1,10,2,0}{4,40,2,0}\n{3,30,1,-1}{5,50,1,-1}"
	assert.Equal(t, tree.levelInfo(), info, fmt.Sprintf("tree.levelInfo() != %s", info))

	//    	    k2(b:1)
	//     	    /    \
	//     k1(b:0)  k4(r:0)
	//   			  / \
	//   	   k3(b:0)  k5(b:0)
	//   	   			  \
	//   	   			  k8(r:-1)
	tree.Insert(k8, v8)
	info = "{2,20,2,1}\n{1,10,2,0}{4,40,1,0}\n{3,30,2,0}{5,50,2,0}\n{8,80,1,-1}"
	assert.Equal(t, tree.levelInfo(), info, fmt.Sprintf("tree.levelInfo() != %s", info))

	assert.Equal(t, tree.Height(), 1, "tree.Height() != 1")
	assert.Equal(t, tree.Search(k2).(int), v2, "tree.Search(k2).(int) != v2")
	assert.Equal(t, tree.Search(k3).(int), v3, "tree.Search(k3).(int) != v3")
	assert.Equal(t, tree.Search(k1).(int), v1, "tree.Search(k3).(int) != v1")
	assert.Nil(t, tree.Search(k7), "tree.Search(k7) != nil")

	//    	    k2(b:1)
	//     	    /    \
	//     k1(b:0)  k4(r:0)
	//   			  / \
	//   	   k3(b:0)  k5(b:0)
	//   	   			  \
	//   	   			  k8(r:-1)
	//   	   			  	/
	//   	   			 k7(r:-1)

	//    	    k2(b:1)
	//     	    /    \
	//     k1(b:0)  k4(r:0)
	//   			  / \
	//   	   k3(b:0)  k7(b:0)
	//   	   			/  \
	//   	   	  k5(r:-1)   k8(r:-1)
	tree.Insert(k7, v7)
	info = "{2,20,2,1}\n{1,10,2,0}{4,40,1,0}\n{3,30,2,0}{7,70,2,0}\n{5,50,1,-1}{8,80,1,-1}"
	assert.Equal(t, tree.levelInfo(), info, fmt.Sprintf("tree.levelInfo() != %s", info))

	//    	    k2(b:1)
	//     	    /    \
	//     k1(b:0)  k4(r:0)
	//   			  / \
	//   	   k3(b:0)  k7(b:0)
	//   	   			/  \
	//   	   	  k5(r:-1)   k8(r:-1)
	//				   \
	//   	   	     k6(r:-1)

	//    	    k2(b:1)
	//     	    /    \
	//     k1(b:0)  k4(r:0)
	//   			  / \
	//   	   k3(b:0)  k7(r:0)
	//   	   			/  \
	//   	   	  k5(b:0)   k8(b:0)
	//				   \
	//   	   	     k6(r:-1)

	//       		 k4(b:1)
	//   		 	/      \
	//   	   k2(r:0)	    k7(r:0)
	//		    /  \		  /    \
	//	 k1(b:0)  k3(b:0) k5(b:0)  k8(b:0)		// k8 为黑兄存在红子
	//					   	  \
	//   	   	  		    k6(r:-1)
	tree.Insert(k6, v6)
	info = "{4,40,2,1}\n{2,20,1,0}{7,70,1,0}\n{1,10,2,0}{3,30,2,0}{5,50,2,0}{8,80,2,0}\n{6,60,1,-1}"
	assert.Equal(t, tree.levelInfo(), info, fmt.Sprintf("tree.levelInfo() != %s", info))
	assert.Equal(t, tree.Height(), 1, "tree.Height() != 1")
	assert.Equal(t, tree.Size(), 8, "tree.Size() != 8")

	//       		 k4(b:1)
	//   		 	/      \
	//   	   k2(r:0)	    k6(r:0)
	//		    /  \		  /    \
	//	 k1(b:0)  k3(b:0) k5(b:0)  k7(b:0)	// k7 为黑兄无红子,且父节点为红节点
	tree.Remove(k8)
	info = "{4,40,2,1}\n{2,20,1,0}{6,60,1,0}\n{1,10,2,0}{3,30,2,0}{5,50,2,0}{7,70,2,0}"
	assert.Equal(t, tree.levelInfo(), info, fmt.Sprintf("tree.levelInfo() != %s", info))

	//       		 k4(b:1)
	//   		 	/      \
	//   	   k2(r:0)	    k6(b:0)
	//		    /  \		  /
	//	 k1(b:0)  k3(b:0) k5(r:0)			// k6 存在红子节点，直接接替
	tree.Remove(k7)
	info = "{4,40,2,1}\n{2,20,1,0}{6,60,2,0}\n{1,10,2,0}{3,30,2,0}{5,50,1,-1}"
	assert.Equal(t, tree.levelInfo(), info, fmt.Sprintf("tree.levelInfo() != %s", info))

	//       		 k4(b:1)
	//   		 	/      \
	//   	   k2(r:0)	    k5(b:0)			// k5 存在红兄弟节点
	//		    /  \
	//	 k1(b:0)  k3(b:0)
	tree.Remove(k6)
	info = "{4,40,2,1}\n{2,20,1,0}{5,50,2,0}\n{1,10,2,0}{3,30,2,0}"
	assert.Equal(t, tree.levelInfo(), info, fmt.Sprintf("tree.levelInfo() != %s", info))

	//       		 k4(b:1)
	//   		 	/      \
	//   	   k2(r:0)	    k5(b:0)			// k5 存在红兄弟节点
	//		    /  \
	//	 k1(b:0)  k3(b:0)

	//       		 k2(b:1)
	//   		 	/      \
	//   	   k1(b:0)	    k4(r:0)
	//					    /     \
	//		  			k3(b:0)   k5(b:0)

	//       		 k2(b:1)
	//   		 	/      \
	//   	   k1(b:0)	    k4(r:0)
	//					    /     \
	//		  			k3(b:0)   k5(b:0)	// 旋转之后 k5 为 黑兄弟 + 红父亲

	//       		 k2(b:1)
	//   		 	/      \
	//   	   k1(b:0)	    k4(r:0)
	//					    /
	//		  			k3(b:0)		// 此时 k3 为黑兄弟+ 红父亲
	tree.Remove(k5)
	info = "{2,20,2,1}\n{1,10,2,0}{4,40,2,0}\n{3,30,1,-1}"
	assert.Equal(t, tree.levelInfo(), info, fmt.Sprintf("tree.levelInfo() != %s", info))

	//       		 k2(b:1)
	//   		 	/      \
	//   	   k1(b:0)	    k4(b:0)
	tree.Remove(k3)
	info = "{2,20,2,1}\n{1,10,2,0}{4,40,2,0}"
	assert.Equal(t, tree.levelInfo(), info, fmt.Sprintf("tree.levelInfo() != %s", info))

	//       		 k2(b:0)		// 合并 k2 与 k4
	//   		 	      \
	//   	            k4(s:0)
	tree.Remove(k1)
	info = "{2,20,2,0}\n{4,40,1,-1}"
	assert.Equal(t, tree.levelInfo(), info, fmt.Sprintf("tree.levelInfo() != %s", info))
}
