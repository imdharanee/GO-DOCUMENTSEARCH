package utils

type Index map[string][]int 

func (idx Index) Add(docs []document) {
   for _,doc:=range docs {
	 for _,token:= analyze(doc.Text) {

		ids:=idx[token]
		if ids!=nil && ids[len(ids)-1]==doc.ID {
			continue
		}
		idx[token]=append(ids,doc.ID)

	 }
   }

}
func Intersection(a []int, b[]int)[]int {

	maxi:=len(a)
	if len(b)>maxi {
		maxi=len(b)
	}
	r:=make([]int,0,maxi)
	var i,j int 
	for i<len(a) && j<len(b) {
		if a[i]<b[j] {
			i++
		} else if(a[i]>b[i]) {
			j++
		
		}else {
			r=append(r,a[i])
			i++
			j++
		}
		
	}
	return r 
}


func (idx Index) Search(text string)[]int {
   var r []int 
   for _,token:=range analyze(text) {
      if ids,ok:=idx[token]; ok {
		if r==nil {
			r=ids 
		}
		else r=Intersection(r,ids)
	  }
	
   }

}