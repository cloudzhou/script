package test_jcseg;

import java.util.Arrays;
import java.util.Comparator;

public class Test {

	public static void main(String[] args) {
		CharTree charTree = new CharTree();
		charTree.appendWord("abc");
		charTree.appendWord("abd");
		System.out.println(charTree.containWord("bcd"));
		System.out.println(charTree.containWord("abc"));
		System.out.println(charTree.containWord("abd"));
	}
	
	public static class CharTree {
		
		public CharNode root = new CharNode(' ', null);
		
		public CharTree() {}
		
		public void appendWord(String word) {
			CharNode currentCharNode = root;
			for(char c : word.toCharArray()) {
				currentCharNode = appendCharNode(currentCharNode, c);
			}
		}
		
		private CharNode appendCharNode(CharNode currentCharNode, char c) {
			CharNode[] charNodes = currentCharNode.getCharNodes();
			if(charNodes == null) {
				charNodes = new CharNode[0];
				currentCharNode.setCharNodes(charNodes);
			}
			CharNode childCharNode = new CharNode(c, null);
			int index =  Arrays.binarySearch(charNodes, childCharNode);
			if(index < 0) {
				CharNode[] newCharNodes =  Arrays.copyOf(charNodes, charNodes.length + 1);
				newCharNodes[newCharNodes.length - 1] = childCharNode;
				Arrays.sort(newCharNodes, new Comparator<CharNode>(){

					@Override
					public int compare(CharNode o1, CharNode o2) {
						if(o1.getC() == o2.getC()) {
							return 0;
						}
						return o1.getC() > o2.getC() ? 1 : -1;
					}});
				charNodes = newCharNodes;
				currentCharNode.setCharNodes(charNodes);
				System.out.println(childCharNode.getC());
				index =  Arrays.binarySearch(charNodes, childCharNode);
			}
			return currentCharNode.getCharNodes()[index];
		}

		public boolean containWord(String word) {
			CharNode currentCharNode = root;
			for(char c : word.toCharArray()) {
				currentCharNode = containCharNode(currentCharNode, c);
				if(currentCharNode == null) {
					return false;
				}
			}
			return true;
		}

		private CharNode containCharNode(CharNode currentCharNode, char c) {
			CharNode[] charNodes =currentCharNode.getCharNodes();
			CharNode childCharNode = new CharNode(c, null);
			int index =  Arrays.binarySearch(charNodes, childCharNode);
			if(index >= 0) {
				return charNodes[index];
			}
			return null;
		}
	}
	
	public static class CharNode implements Comparable<CharNode> {
		
		public char c;
		public CharNode[] charNodes;
		
		public CharNode(char c, CharNode[] charNodes) {
			this.c = c;
			this.charNodes = charNodes;
		}
		
		public char getC() {
			return c;
		}
		public void setC(char c) {
			this.c = c;
		}
		public CharNode[] getCharNodes() {
			return charNodes;
		}
		public void setCharNodes(CharNode[] charNodes) {
			this.charNodes = charNodes;
		}

		@Override
		public int compareTo(CharNode o) {
			if(this.getC() == o.getC()) {
				return 0;
			}
			return this.getC() > o.getC() ? 1 : -1;
		}
		
	}
}

