package datastruct

// Hash - Rolling Hash
func Hash(s string) int {
	h := 0
	A := 256 // A는 곱하고 S와 더하니까 255보다 큰 값을 설정하는게 좋음
	// * 왜 255보다 큰 값인가? S는 문자고 문자는 ASCII CODE, 0~255의 값을 가지고 있음
 	B := 3571 // 나머지 연산자는 소수로 하는게 좋다 => 값의 분포도가 넓게 퍼지기 때문
	for i := 0; i < len(s); i++ {
		// H0 = (S0) % B -> H1 = (H0 * A + S) % B -> H2 = (H1 * A + S) % B ... 
		h = (h*A + int(s[i])) % B
	}
	return h
}
/*
 문자열 입력은 무한대, 출력은 0~3570 까지 -> 손실 압축
 Hash 충돌로 같은 값이 나올 수 있음, 어떻게 해결하는가?
 => 또 다른 배열을 추가하여 같은 Hash 값을 갖는 Key들을 저장해 충돌 방지
*/

type keyValue struct {
	key   string
	value string
}

// CreateMap Map 주소를 return함
func CreateMap() *Map {
	return &Map{}
}

// Map ... 
type Map struct {
	// [3571] = keyArray / [] = 중복된 keyValue들 저장
	keyArray [3571][]keyValue
}

// Add ...
func (m *Map) Add(key, value string) {
	h := Hash(key) // key를 Hash로 돌려서 나온 값을
	m.keyArray[h] = append(m.keyArray[h], keyValue{key, value}) // keyArray에 append
}

// Get - 위에서 Add된걸 가져옴
func (m *Map) Get(key string) string {
	h := Hash(key)
	for i := 0; i < len(m.keyArray[h]); i++ {
		if m.keyArray[h][i].key == key { 
			return m.keyArray[h][i].value // key가 같으면 value return
		}
	}
	return "" // 없으면 빈 문자열 return
}