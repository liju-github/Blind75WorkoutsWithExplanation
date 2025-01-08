func reorderList(head *ListNode) {
    if head == nil || head.Next == nil {
        return
    }

    //find middle
    //reverse second half
    //merge alternatively
    fast,slow := head,head

    for fast!=nil && fast.Next !=nil{
        fast= fast.Next.Next
        slow = slow.Next
    }

    current := slow.Next
    slow.Next = nil // remove connection to sec half
    var prev *ListNode
    for current!=nil{
        next := current.Next
        current.Next = prev
        prev = current
        current = next
    }

    first,second := head,prev

    for second!=nil{
        tmp1,tmp2:= first.Next,second.Next
        
        first.Next = second
        second.Next = tmp1

        first = tmp1
        second = tmp2
    }

}














