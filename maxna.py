def max_score(t, test_cases):
    results = []

    for n, array in test_cases:
        max_score = 0

        for start in range(2):
            selected = []
            for i in range(start, n, 2):
                selected.append(array[i])

            if selected:
                current_max = max(selected)
                current_min = min(selected)
                count_red = len(selected)
                score = current_max + current_min + count_red
                max_score = max(max_score, score)

                # Debugging output
                print(f"Start from {start}: Selected = {selected}, Max = {current_max}, Min = {current_min}, Count = {count_red}, Score = {score}")

        results.append(max_score)

    return results

# Example usage
t = 4
test_cases = [
    (3, [5, 4, 5]),
    (3, [4, 5, 4]),
    (10, [3, 3, 3, 3, 4, 1, 2, 3, 5, 4]),
    (10, [17, 89, 92, 42, 29, 41, 92, 14, 70, 45]),
]

output = max_score(t, test_cases)
for score in output:
    print(score)
