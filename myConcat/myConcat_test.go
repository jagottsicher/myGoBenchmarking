package myConcat

import (
	"fmt"
	"strings"
	"testing"
)

func TestConcat(t *testing.T) {
      kafkaOuttake := "One morning, when ..."
      sliceOfStringOfkafkaOuttake := strings.Split(kafkaOuttake, " ")
      kafkaOuttake = Concat(sliceOfStringOfkafkaOuttake)
      if kafkaOuttake != "One morning, when ..." {
              t.Error("Expected:", "One morning, when ...", "got:", kafkaOuttake)
      }
}

func TestJoin(t *testing.T) {
	kafkaOuttake := "One morning, when ..."
	sliceOfStringOfkafkaOuttake := strings.Split(kafkaOuttake, " ")
	kafkaOuttake = Join(sliceOfStringOfkafkaOuttake, " ")
	if kafkaOuttake != "One morning, when ..." {
		t.Error("Expected:", "One morning, when ...", "got:", kafkaOuttake)
	}
}


 func ExampleConcat() {
      sliceOfStringOfkafkaOuttake := []string{"One", "morning,", "when", "..."}
      fmt.Println(Concat(sliceOfStringOfkafkaOuttake))
      // Output:
      // One morning, when ...
 }

const kafka = `"One morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin. He lay on his armour-like back, and if he lifted his head a little he could see his brown belly, slightly domed and divided by arches into stiff sections. The bedding was hardly able to cover it and seemed ready to slide off any moment. His many legs, pitifully thin compared with the size of the rest of him, waved about helplessly as he looked. "What's happened to me?" he thought. It wasn't a dream. His room, a proper human room although a little too small, lay peacefully between its four familiar walls. A collection of textile samples lay spread out on the table - Samsa was a travelling salesman - and above it there hung a picture that he had recently cut out of an illustrated magazine and housed in a nice, gilded frame. It showed a lady fitted out with a fur hat and fur boa who sat upright, raising a heavy fur muff that covered the whole of her lower arm towards the viewer. Gregor then turned to look out the window at the dull weather."`

var sliceStringTemp []string

func BenchmarkConcat(b *testing.B) {
	sliceStringTemp = strings.Split(kafka, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Concat(sliceStringTemp)
	}
}

func BenchmarkJoin(b *testing.B) {
	sliceStringTemp = strings.Split(kafka, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(sliceStringTemp, " ")
	}
}
