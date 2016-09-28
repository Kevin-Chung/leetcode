import(
    "strconv"   
)    
func twoSum(nums []int, target int) []int {
    var m = make(map[int]int);
    var compliment int;
    for i:= 0 ; i < len(nums); i++{
        compliment = target - nums[i];
        _, key := m[compliment];
        if(key){
            return []int{m[compliment],i};
        }
        m[nums[i]] = i;
    }
    return []int{};   
}